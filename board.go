package main

import (
	"fmt"
)

const (
	Empty = 1 << iota
	Black
	White
)

type board struct {
	rows [][]int
}

func (b *board) printboard() {
	for _, row := range b.rows {
		for _, square := range row {
			var ch string
			switch square {
			case White:
				ch = "●"
			case Black:
				ch = "○"
			default:
				ch = "."
			}
			fmt.Printf("%s ", ch)
		}
		fmt.Println("")
	}
}
