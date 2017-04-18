package main

import (
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

func (pb *PlyBoard) deepSearch(b *Board, myColour Square, depth int) *Move {
	bcp := b.copy()
	// TODO return result on a chan
	moves := pb.findAllMoves(bcp, myColour)
	if len(moves) == 0 {
		return nil
	}
	if len(moves) == 1 {
		return &moves[0]
	}
	sort.Sort(sort.Reverse(ByScore(moves)))
	for i := 0; i < len(moves); i++ {
		//pos := scores[i].pos
		// opp has valid?
		// swap piece
		// find all moves
		// go deeper
		pb.xx(bcp, myColour, 2)
		// else have i got valid?
	}
	bestMove := &moves[0]
	return bestMove
}

func (pb *PlyBoard) findAllMoves(b *Board, myColour Square) []Move {
	var pos Position
	moves := []Move{}
	validMoves := b.findAllValidMoves(myColour)
	for i := 0; i < len(validMoves); i++ {
		bcp := b.copy()
		bcp.playMove(&validMoves[i], myColour)
		score := pb.evaluationFunction(bcp.rows, myColour)
		moves = append(moves, Move{pos: &pos, score: score})
	}
	return moves
}

func (pb *PlyBoard) xx(b *Board, myPiece Square, depth int) {
	if depth -= 1; depth == 0 {
		// TODO do evaluation
		return
	}

}

func (pb *PlyBoard) oppColour(myColour Square) Square {
	opp := Black
	if myColour == Black {
		opp = White
	}
	return opp
}
