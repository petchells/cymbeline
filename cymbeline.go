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

func main() {
	rand.Seed(time.Now().UnixNano())
	b := playGame()
	if b != nil {
		whiteCount, blackCount := b.countPieces()
		fmt.Println("White: %d", whiteCount)
		fmt.Println("Black: %d", blackCount)
	} else {
		fmt.Println("User quit")
	}
}

func playGame() *Board {
	b := newBoard()
	p1 := Black
	p2 := White
	movePossible := false
	b.printboard()

	for {
		if b.hasValidMove(p1) {
			movePossible = true
			pos := b.findBestMove(p1)
			b.playMove(pos, p1)
			fmt.Println("I've played " + pos.AsString())
		} else {
			fmt.Println("I can't go")
			if !movePossible {
				return b
			}
			movePossible = false
		}
		if b.hasValidMove(p2) {
			movePossible = true
			b.printboard()
			userMove := getHumanMove(b, p2)
			if userMove == nil {
				return nil
			}
			b.playMove(userMove, p2)
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
