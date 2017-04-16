package main

import (
	"sort"
)

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

func deepSearch(b *Board, myColour Square, depth int) *Score {
	scores := findAllMoves(b, myColour)
	if len(scores) == 0 {
		return nil
	}
	sort.Sort(ByScore(scores))
	for i := 0; i < len(scores); i++ {
		bcp := b.copy()
		//pos := scores[i].pos
		xx(bcp, myColour, 2)
		// opp has valid?
		// swap piece
		// find all moves
		// go deeper
		// else have i got valid?
	}
	bestScore := Score{}
	return &bestScore
}

func findAllMoves(b *Board, myColour Square) []Score {
	var pos Position
	scores := []Score{}
	validMoves := b.findAllValidMoves(myColour)
	for i := 0; i < len(validMoves); i++ {
		bcp := b.copy()
		bcp.playMove(&validMoves[i], myColour)
		score := dynamic_heuristic_evaluation_function_alt(bcp.rows, myColour)
		scores = append(scores, Score{pos: &pos, score: score})
	}
	return scores
}

func xx(b *Board, myPiece Square, depth int) {

}
func oppColour(myColour Square) Square {
	opp := Black
	if myColour == Black {
		opp = White
	}
	return opp
}
