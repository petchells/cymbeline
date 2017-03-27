package main

import (
	"errors"
	"math/rand"
)

/**
 * Find the best move for a given board position
 */

func findBestMove(b *Board, myPiece Square) (*Position, error) {
	// Iterate over the top level moves
	for {
		pos := &Position{
			x: int8(rand.Intn(len(b.rows))),
			y: int8(rand.Intn(len(b.rows)))}
		counter := 0
		if playMove(b, pos, myPiece) {
			return pos, nil
		} else {
			counter++
			if counter > 100 {
				return nil, errors.New("I give up")
			}
		}
	}
	return nil, errors.New("No more moves")
}
func isOnBoard(b *Board, p *Position) bool {
	return p.x >= 0 && p.y >= 0 &&
		p.x < int8(len(b.rows)) && p.y < int8(len(b.rows))
}
func scanDiagonal(b *Board, p *Position, myPiece Square, xinc int8, yinc int8) []Position {
	turned := make([]Position, 0, len(b.rows))
	foundOpp := false
	nextPos := Position{x: p.x + xinc, y: p.y + yinc}
	for isOnBoard(b, &nextPos) {
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
func scanAllDiagonals(b *Board, p *Position, myPiece Square) []Position {
	turned := make([]Position, 0, len(b.rows)*4)
	for i := int8(-1); i <= 1; i++ {
		for j := int8(-1); j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			positions := scanDiagonal(b, p, myPiece, i, j)
			if len(positions) > 0 {
				turned = append(turned, positions...)
			}
		}
	}
	return turned
}
func findTurned(b *Board, p *Position, myPiece Square) []Position {
	if !isOnBoard(b, p) || b.getSquare(p) != Empty {
		return []Position{}
	}
	// look along all diagonals for turned pieces
	return scanAllDiagonals(b, p, myPiece)
}
func playMove(b *Board, p *Position, myPiece Square) bool {
	turned := findTurned(b, p, myPiece)
	if len(turned) == 0 {
		return false
	}
	b.setPiece(p, myPiece)
	for _, pos := range turned {
		b.setPiece(&pos, myPiece)
	}
	return true
}
