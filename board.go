package main

import (
	"fmt"
	"regexp"
)

type Square int8

const (
	Empty Square = 1 << iota
	Black
	White
)

type Board struct {
	rows [][]Square
}

type Position struct {
	x, y int8
}

var validPositionString = regexp.MustCompile(`^[A-Z][1-8]$`)

func positionFromString(s string) *Position {
	// Use raw strings to avoid having to quote the backslashes.
	if validPositionString.MatchString(s) {
		x := 56 - s[1]
		y := s[0] - 65
		return &Position{x: int8(x), y: int8(y)}
	}
	return nil
}

func (p *Position) AsString() string {
	return fmt.Sprintf("%c%c", p.y+65, 56-p.x)
}

func newBoard() *Board {
	// The number of rows and columns doesn't need to be 8 (the code uses
	// len(rows) when iterating). It *does* have to be square, though.
	rows := [][]Square{
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Black, White, Empty, Empty, Empty},
		{Empty, Empty, Empty, White, Black, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}
	return &Board{rows: rows}
}
func (b *Board) getSquare(p *Position) Square {
	return b.rows[p.x][p.y]
}
func (b *Board) setPiece(pos *Position, piece Square) {
	b.rows[pos.x][pos.y] = piece
}
func (b *Board) printboard() {
	fmt.Print(" ")
	for i, _ := range b.rows[0] {
		fmt.Printf(" %c", i+65)
	}
	fmt.Println("")
	for i, row := range b.rows {
		fmt.Printf("%c ", 56-i)
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
		fmt.Printf("%c\n", 56-i)
	}
	fmt.Print(" ")
	for i, _ := range b.rows[0] {
		fmt.Printf(" %c", i+65)
	}
	fmt.Println("")
}
