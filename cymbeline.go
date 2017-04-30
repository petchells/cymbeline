// main project cymbeline.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type MoveStrategy func(*Board, Square) *Position

func main() {
	rand.Seed(time.Now().UnixNano())
	webflag := flag.Bool("web", false, "enable web server on port 8088")
	termflag := flag.Bool("term", false, "play a game in the terminal")
	soloflag := flag.Bool("solo", false, "plays with itself")
	flag.Parse()
	if *webflag {
		go serve()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Server running on http://localhost:8088/")
		fmt.Println("Press Enter to quit")
		reader.ReadString('\n')
	} else if *termflag {
		humanVsMachine()
	} else if *soloflag {
		fmt.Println("Optimus Prime is playing 50 games against Wall-E...")
		machineVsMachine(25)
	}
}

func humanVsMachine() {
	b := playGame(optimusPrime, human)
	if b != nil {
		blackCount, whiteCount := b.countPieces()
		fmt.Println("Black: ", blackCount)
		fmt.Println("White: ", whiteCount)
	} else {
		fmt.Println("Quit")
	}
}
func machineVsMachine(nrRounds int) {
	var b *Board
	opCnt, waCnt := 0, 0
	for n := 0; n < nrRounds; n++ {
		b = playGame(optimusPrime, walle)
		if b != nil {
			blackCount, whiteCount := b.countPieces()
			if blackCount > whiteCount {
				opCnt += 1
			} else if whiteCount > blackCount {
				waCnt += 1
			}
		}
		b = playGame(walle, optimusPrime)
		if b != nil {
			blackCount, whiteCount := b.countPieces()
			if blackCount > whiteCount {
				waCnt += 1
			} else if whiteCount > blackCount {
				opCnt += 1
			}
		}
	}
	fmt.Println("Optimus Prime: ", opCnt)
	fmt.Println("Wall-E: ", waCnt)
}

func playGame(p1Mover MoveStrategy, p2Mover MoveStrategy) *Board {
	b := newBoard()
	movePossible := false
	for {
		if b.hasValidMove(Black) {
			movePossible = true
			p1Move := p1Mover(b, Black)
			if p1Move == nil {
				return nil
			}
			b.playMove(p1Move, Black)
		} else {
			//fmt.Println("No moves are possible for black")
			if !movePossible {
				return b
			}
			movePossible = false
		}
		if b.hasValidMove(White) {
			movePossible = true
			p2Move := p2Mover(b, White)
			if p2Move == nil {
				return nil
			}
			b.playMove(p2Move, White)
		} else {
			//fmt.Println("No moves are possible for white")
			if !movePossible {
				return b
			}
			movePossible = false
		}
	}
	return b
}
func optimusPrime(b *Board, colour Square) *Position {
	//move := b.findBestMove(colour)
	pb := PlyBoard{evaluationFunction: dynamic_heuristic_evaluation_function}
	move := pb.deepSearch(b, colour)
	return move.pos
}
func walle(b *Board, colour Square) *Position {
	pb := PlyBoard{evaluationFunction: dynamic_heuristic_evaluation_function_alt}
	move := pb.findBestMoveDither(b, colour)
	return move.pos
}
func human(b *Board, colour Square) *Position {
	b.printboard()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter move: ")
		text, _ := reader.ReadString('\n')
		move := positionFromString(strings.ToUpper(strings.TrimSpace(text)))
		if move == nil {
			return nil
		}
		if b.isValidMove(move, colour) {
			return move
		}
	}
	return nil
}
