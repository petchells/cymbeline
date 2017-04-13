package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
)

type Square int

const (
	Empty Square = 1 << iota
	Black
	White
)

type Board struct {
	rows [8][8]Square
}

func (b *Board) copy() *Board {
	cp := [8][8]Square{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cp[i][j] = b.rows[i][j]
		}
	}
	return &Board{cp}
}

func (b *Board) copyFrom(from *Board) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.rows[i][j] = from.rows[i][j]
		}
	}
}

type Position struct {
	x, y int
}

var validPositionString = regexp.MustCompile(`^[A-Z][1-8]$`)
var validPositionStringRev = regexp.MustCompile(`^[1-8][A-Z]$`)

func enumToColour(s Square) string {
	switch s {
	case Black:
		return "black"
	case White:
		return "white"
	default:
		return ""
	}
}
func positionFromString(s string) *Position {
	// Use raw strings to avoid having to quote the backslashes.
	if validPositionString.MatchString(s) {
		x := s[1] - 49
		y := s[0] - 65
		return &Position{x: int(x), y: int(y)}
	} else if validPositionStringRev.MatchString(s) {
		x := s[0] - 49
		y := s[1] - 65
		return &Position{x: int(x), y: int(y)}
	}
	return nil
}
func (p *Position) AsString() string {
	return fmt.Sprintf("%c%c", 65+p.y, 49+p.x)
}
func newBoard() *Board {
	// The number of rows and columns doesn't need to be 8 (the code uses
	// len(rows) when iterating). It *does* have to be square, though.
	var rows [8][8]Square
	rows = [8][8]Square{
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
		fmt.Printf(" %c", 65+i)
	}
	fmt.Println("")
	for i, row := range b.rows {
		fmt.Printf("%c ", 49+i)
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
		fmt.Printf("%c\n", 49+i)
	}
	fmt.Print(" ")
	for i, _ := range b.rows[0] {
		fmt.Printf(" %c", 65+i)
	}
	fmt.Println("")
}
func (b *Board) hasValidMove(myPiece Square) bool {
	oppPiece := Black
	if myPiece == Black {
		oppPiece = White
	}
	return num_valid_moves(myPiece, oppPiece, b.rows) > 0
}

type ByScore []Score

func (moves ByScore) Len() int {
	return len(moves)
}
func (moves ByScore) Swap(i, j int) {
	moves[i], moves[j] = moves[j], moves[i]
}
func (moves ByScore) Less(i, j int) bool {
	return moves[j].score < moves[i].score
}
func (b *Board) countPieces() (int, int) {
	wht, blk := 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.rows[i][j] == White {
				wht += 1
			} else if b.rows[i][j] == Black {
				blk += 1
			}
		}
	}
	return blk, wht
}

/**
 * Find the best move for a given board position
 */
func (b *Board) findBestMove(myPiece Square) *Position {
	// Iterate over the top level moves
	validMoves := []Score{}
	for i := 0; i < len(b.rows[0]); i++ {
		for j := 0; j < len(b.rows); j++ {
			pos := &Position{
				x: i,
				y: j}
			if b.isValidMove(pos, myPiece) {
				bcp := b.copy()
				bcp.playMove(pos, myPiece)
				score := dynamic_heuristic_evaluation_function(bcp.rows, myPiece)
				validMoves = append(validMoves, Score{pos: pos, score: score})
			}
		}
	}
	if len(validMoves) == 0 {
		return nil
	}
	sort.Sort(ByScore(validMoves))
	move := validMoves[0]
	return move.pos
}
func (b *Board) findBestMoveAlt(myPiece Square) *Position {
	// Iterate over the top level moves
	validMoves := []Score{}
	totalScores := 0.0
	for i := 0; i < len(b.rows[0]); i++ {
		for j := 0; j < len(b.rows); j++ {
			pos := &Position{
				x: i,
				y: j}
			if b.isValidMove(pos, myPiece) {
				bcp := b.copy()
				bcp.playMove(pos, myPiece)
				score := dynamic_heuristic_evaluation_function_alt(bcp.rows, myPiece)
				totalScores += score
				validMoves = append(validMoves, Score{pos: pos, score: score})
			}
		}
	}
	if len(validMoves) == 0 {
		return nil
	}
	sort.Sort(ByScore(validMoves))
	pick := rand.Float64() * totalScores
	for i := 0; i < len(validMoves); i++ {
		totalScores -= validMoves[i].score
		if pick >= totalScores {
			return validMoves[i].pos
		}
	}
	var move Score
	if l := len(validMoves) / 2; l == 0 {
		move = validMoves[0]
	} else {
		move = validMoves[rand.Intn(l)]
	}
	return move.pos
}
func (b *Board) isOnBoard(p *Position) bool {
	return p.x >= 0 && p.y >= 0 && p.x < 8 && p.y < 8
}
func (b *Board) scanDiagonal(p *Position, myPiece Square, xinc int, yinc int) []Position {
	turned := make([]Position, 0, 8)
	foundOpp := false
	nextPos := Position{x: p.x + xinc, y: p.y + yinc}
	for b.isOnBoard(&nextPos) {
		if b.getSquare(&nextPos) == Empty {
			// none of our guys in that line
			break
		}
		if b.getSquare(&nextPos) == myPiece {
			if !foundOpp {
				// found our piece but no opponents - nothing turned
				break
			}
			// found our piece at the end of a line of oppenents
			return turned
		}
		// found a line of opponents. build stack and iterate
		turned = append(turned, nextPos)
		foundOpp = foundOpp || true
		nextPos = Position{
			x: nextPos.x + xinc,
			y: nextPos.y + yinc}
	}
	return []Position{}
}
func (b *Board) findAllValidMoves(myPiece Square) []Position {
	var p Position
	list := []Position{}
	for i := 0; i < len(b.rows[0]); i++ {
		for j := 0; j < len(b.rows); j++ {
			p.x, p.y = i, j
			if b.isValidMove(&p, myPiece) {
				list = append(list, p)
				p = Position{}
			}
		}
	}
	return list
}
func (b *Board) isValidMove(p *Position, myPiece Square) bool {
	if !b.isOnBoard(p) || b.getSquare(p) != Empty {
		return false
	}
	opp := Black
	if myPiece == Black {
		opp = White
	}
	return isLegalMove(myPiece, opp, b.rows, p.x, p.y)

	//	for i := int8(-1); i <= 1; i++ {
	//		for j := int8(-1); j <= 1; j++ {
	//			if i == 0 && j == 0 {
	//				continue
	//			}
	//			positions := b.scanDiagonal(p, myPiece, i, j)
	//			if len(positions) > 0 {
	//				return true
	//			}
	//		}
	//	}
	//return false
}
func (b *Board) findTurned(p *Position, myPiece Square) []Position {
	if p == nil || !b.isOnBoard(p) || b.getSquare(p) != Empty {
		return []Position{}
	}
	// look along all diagonals for turned pieces
	turned := make([]Position, 0, len(b.rows)*4)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			positions := b.scanDiagonal(p, myPiece, i, j)
			if len(positions) > 0 {
				turned = append(turned, positions...)
			}
		}
	}
	return turned
}
func (b *Board) playMove(p *Position, myPiece Square) bool {
	turned := b.findTurned(p, myPiece)
	if len(turned) == 0 {
		return false
	}
	b.setPiece(p, myPiece)
	for _, pos := range turned {
		b.setPiece(&pos, myPiece)
	}
	return true
}
