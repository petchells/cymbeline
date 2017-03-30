package main

var posWeights = [][]int{
	{20, -3, 11, 8, 8, 11, -3, 20},
	{-3, -7, -4, 1, 1, -4, -7, -3},
	{11, -4, 2, 2, 2, 2, -4, 11},
	{8, 1, 2, -3, -3, 2, 1, 8},
	{8, 1, 2, -3, -3, 2, 1, 8},
	{11, -4, 2, 2, 2, 2, -4, 11},
	{-3, -4, -1, -1, -1, -1, -4, -3},
	{20, -3, 11, 8, 8, 11, -3, 20}}

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
