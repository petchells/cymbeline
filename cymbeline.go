// main project cymbeline.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	b := newBoard()
	b.printboard()
	pos, err := findBestMove(b, White)
	if err == nil {
		fmt.Println("I've played " + pos.AsString())
		b.printboard()
	} else {
		fmt.Printf("%q", err)
	}
}
