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

type MoveStrategy func(*Board, Square) *Position

func main() {
	rand.Seed(time.Now().UnixNano())
	b := playGame(optimusPrime, human)
	if b != nil {
		blackCount, whiteCount := b.countPieces()
		fmt.Println("Black: ", blackCount)
		fmt.Println("White: ", whiteCount)
	} else {
		fmt.Println("User quit")
	}
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
			fmt.Println("No moves are possible for black")
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
			fmt.Println("No moves are possible for white")
			if !movePossible {
				return b
			}
			movePossible = false
		}
	}
	return b
}
func optimusPrime(b *Board, colour Square) *Position {
	move := b.findBestMove(colour)
	fmt.Println("I played: " + move.AsString())
	return move
}
func walle(b *Board, colour Square) *Position {
	move := b.findBestMoveAlt(colour)
	fmt.Println("I2 played: " + move.AsString())
	return move
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
