package main

import (
	"math"
	"sort"
)

type ByScore []Move

func (moves ByScore) Len() int {
	return len(moves)
}

func (moves ByScore) Swap(i, j int) {
	moves[i], moves[j] = moves[j], moves[i]
}

func (moves ByScore) Less(i, j int) bool {
	return moves[j].score < moves[i].score
}

type PlyBoard struct {
	evaluationFunction func([8][8]Square, Square) float64
}

func (pb *PlyBoard) deepSearch(b *Board, myColour Square, depth int) Move {
	best := Move{score: math.Inf(-1)}
	recursiveSearch := func(b *Board, colour Square, depth int) Move {
		pos := Position{}
		bcp := &Board{}
		oppColour := pb.oppColour(colour)
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				pos.x, pos.y = i, j
				if b.isValidMove(pos, colour) {
					bcp.copyFrom(b)
					b.playMove(pos, colour)
					if depth == 0 {
						score := pb.evaluationFunction(bcp.rows, colour)
						return Move{pos: &pos, score: score}
					}
					// switch sides and go deeper
					m := pb.recursiveSearch(b, oppColour, depth-1)
					if m.score > best.score {
						best
					}
				}
			}
		}
		return nil
	}
	return best
}

func (pb *PlyBoard) findAllMoves(b *Board, myColour Square) []Move {
	var pos Position
	moves := []Move{}
	positions := b.findAllValidMoves(myColour)
	bcp := &Board{}
	for _, move := range positions {
		bcp.copyFrom(b)
		bcp.playMove(&move, myColour)
		score := pb.evaluationFunction(bcp.rows, myColour)
		moves = append(moves, Move{pos: &pos, score: score})
	}
	return moves
}

func (pb *PlyBoard) oppColour(myColour Square) Square {
	opp := Black
	if myColour == Black {
		opp = White
	}
	return opp
}
