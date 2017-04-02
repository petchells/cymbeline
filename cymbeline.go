// main project cymbeline.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Mover func(b *Board, colour *Square) *Position

func main() {
	rand.Seed(time.Now().UnixNano())
	b := playGame(getHumanMove, getComputerMove1)
	if b != nil {
		blackCount, whiteCount := b.countPieces()
		fmt.Println("Black: %d", blackCount)
		fmt.Println("White: %d", whiteCount)
	} else {
		fmt.Println("User quit")
	}
}

func playGame(p1Mover Mover, p2Mover Mover) *Board {
	b := newBoard()
	p1 := Black
	p2 := White
	movePossible := false

	for {
		if b.hasValidMove(p1) {
			movePossible = true
			p1Move := getComputerMove1(b, p1)
			b.playMove(p1Move, p1)
			fmt.Println("I've played " + p1Move.AsString())
		} else {
			fmt.Println("I can't go")
			if !movePossible {
				return b
			}
			movePossible = false
		}
		if b.hasValidMove(p2) {
			movePossible = true
			p2Move := getHumanMove(b, p2)
			if p2Move == nil {
				return nil
			}
			b.playMove(p2Move, p2)
		} else {
			fmt.Println("You can't go")
			if !movePossible {
				return b
			}
			movePossible = false
		}
	}
	return b
}
func getComputerMove1(b *Board, colour Square) *Position {
	return b.findBestMove(colour)
}
func getComputerMove2(b *Board, colour Square) *Position {
	return b.findBestMoveAlt(colour)
}
func getHumanMove(b *Board, colour Square) *Position {
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
