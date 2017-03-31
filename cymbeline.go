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

func playGame() *Board {
	b := newBoard()
	userPiece := Black
	myPiece := White
quit:
	for {
		b.printboard()
		m := Mover{b}
		pos, err := m.findBestMove(myPiece)
		m.playMove(pos, myPiece)
		if err == nil {
			fmt.Println("I've played " + pos.AsString())
			b.printboard()
		} else {
			fmt.Printf("%q", err)
		}
		for {
			move := strings.ToUpper(strings.TrimSpace(getInput()))
			userMove := positionFromString(move)
			if userMove == nil {
				fmt.Println("Quitting")
				break quit
			}
			if m.isValidMove(userMove, userPiece) {
				m.playMove(userMove, userPiece)
				break
			}
		}
	}
	return b
}
func main() {
	rand.Seed(time.Now().UnixNano())
	playGame()
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move: ")
	text, _ := reader.ReadString('\n')
	return text
}
