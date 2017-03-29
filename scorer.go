package main

type ScoreParams struct {
	// multiplier of piece count for each move
	pieceCountWeights []int
	// board position scores
	positionWeights [][]int
	recurseDepth    int
}

type Scorer struct {
	b       *Board
	myPiece Square
	params  *ScoreParams
}

func (s *Scorer) calculateScore() int {
	return 1
}
