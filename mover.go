package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Mover struct {
	b *Board
}
type Score struct {
	pos   *Position
	score float64
}

/**
 * Find the best move for a given board position
 */
func (m *Mover) findBestMove(myPiece Square) (*Position, error) {
	// Iterate over the top level moves
	validMoves := []Score{}
	for i := 0; i < len(m.b.rows[0]); i++ {
		for j := 0; j < len(m.b.rows); j++ {
			pos := &Position{
				x: int8(i),
				y: int8(j)}
			if m.isValidMove(pos, myPiece) {
				m.playMove(pos, myPiece)
				m.b.printboard()
				score := dynamic_heuristic_evaluation_function(m.b.rows, myPiece)
				validMoves = append(validMoves, Score{pos: pos, score: score})
				fmt.Printf("Score: %f, Position: %s\n", score, pos.AsString())
			}
		}
	}
	if len(validMoves) == 0 {
		return nil, errors.New("No moves possible")
	}
	move := validMoves[rand.Intn(len(validMoves))]
	return move.pos, nil
}
func (m *Mover) findBestMoveAlt(myPiece Square) (*Position, error) {
	// Iterate over the top level moves
	validMoves := make([]Score, 0, 30)
	pos := &Position{}
	for i := 0; i < len(m.b.rows[0]); i++ {
		for j := 0; j < len(m.b.rows); j++ {
			pos.x = int8(i)
			pos.y = int8(j)
			if m.isValidMove(pos, myPiece) {
				m.playMove(pos, myPiece)
				score := dynamic_heuristic_evaluation_function(m.b.rows, myPiece)
				validMoves = append(validMoves, Score{pos: pos, score: score})
				fmt.Printf("Score: %f, Position: %s\n", score, pos.AsString())
			}
		}
	}
	if len(validMoves) == 0 {
		return nil, errors.New("No moves possible")
	}
	move := validMoves[rand.Intn(len(validMoves))]
	return move.pos, nil
}
func (m *Mover) isOnBoard(p *Position) bool {
	return p.x >= 0 && p.y >= 0 &&
		p.x < int8(len(m.b.rows[0])) && p.y < int8(len(m.b.rows))
}
func (m *Mover) scanDiagonal(p *Position, myPiece Square, xinc int8, yinc int8) []Position {
	turned := make([]Position, 0, len(m.b.rows))
	foundOpp := false
	nextPos := Position{x: p.x + xinc, y: p.y + yinc}
	for m.isOnBoard(&nextPos) {
		if m.b.getSquare(&nextPos) == Empty {
			// none of our guys in that line
			break
		}
		if m.b.getSquare(&nextPos) == myPiece {
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
func (m *Mover) isValidMove(p *Position, myPiece Square) bool {
	if !m.isOnBoard(p) || m.b.getSquare(p) != Empty {
		return false
	}
	for i := int8(-1); i <= 1; i++ {
		for j := int8(-1); j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			positions := m.scanDiagonal(p, myPiece, i, j)
			if len(positions) > 0 {
				return true
			}
		}
	}
	return false
}
func (m *Mover) findTurned(p *Position, myPiece Square) []Position {
	if !m.isOnBoard(p) || m.b.getSquare(p) != Empty {
		return []Position{}
	}
	// look along all diagonals for turned pieces
	turned := make([]Position, 0, len(m.b.rows)*4)
	for i := int8(-1); i <= 1; i++ {
		for j := int8(-1); j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			positions := m.scanDiagonal(p, myPiece, i, j)
			if len(positions) > 0 {
				turned = append(turned, positions...)
			}
		}
	}
	return turned
}
func (m *Mover) playMove(p *Position, myPiece Square) bool {
	turned := m.findTurned(p, myPiece)
	if len(turned) == 0 {
		return false
	}
	m.b.setPiece(p, myPiece)
	for _, pos := range turned {
		m.b.setPiece(&pos, myPiece)
	}
	return true
}
