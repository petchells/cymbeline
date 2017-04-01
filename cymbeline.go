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
	playGame()
}

func playGame() *Board {
	b := newBoard()
	p1 := Black
	p2 := White
	movePossible := false
	b.printboard()
quit:
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
			for {
				b.printboard()
				userMove := getHumanMove()
				if userMove == nil {
					fmt.Println("Quitting")
					break quit
				}
				if b.isValidMove(userMove, p2) {
					b.playMove(userMove, p2)
					break
				}
			}
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

func getHumanMove() *Position {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move: ")
	text, _ := reader.ReadString('\n')
	return positionFromString(strings.ToUpper(strings.TrimSpace(text)))
}
