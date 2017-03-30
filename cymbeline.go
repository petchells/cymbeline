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
	b := newBoard()
	userPiece := Black
	myPiece := White
quit:
	for {
		sp1 := scoreParams1()
		fmt.Printf("params %q\n", sp1.pieceCountWeights)
		// s := Scorer{b: b, myPiece: myPiece, params: &sp1}
		// fmt.Printf("Scorer %q\n", s)
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
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move: ")
	text, _ := reader.ReadString('\n')
	return text
}
