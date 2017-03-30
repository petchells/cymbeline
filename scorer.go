package main

var posWeights = [][]int{
	{4, -3, 2, 2, 2, 2, -3, 4},
	{-3, -4, -1, -1, -1, -1, -4, -3},
	{2, -1, 1, 0, 0, 1, -1, 2},
	{2, -1, 0, 1, 1, 0, -1, 2},
	{2, -1, 0, 1, 1, 0, -1, 2},
	{2, -1, 1, 0, 0, 1, -1, 2},
	{-3, -4, -1, -1, -1, -1, -4, -3},
	{4, -3, 2, 2, 2, 2, -3, 4}}

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
