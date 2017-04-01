package main

import (
	// "fmt"
	"math"
)

var posWeights = [][]int{
	{20, -3, 11, 8, 8, 11, -3, 20},
	{-3, -7, -4, 1, 1, -4, -7, -3},
	{11, -4, 2, 2, 2, 2, -4, 11},
	{8, 1, 2, -3, -3, 2, 1, 8},
	{8, 1, 2, -3, -3, 2, 1, 8},
	{11, -4, 2, 2, 2, 2, -4, 11},
	{-3, -4, -1, -1, -1, -1, -4, -3},
	{20, -3, 11, 8, 8, 11, -3, 20}}

type Score struct {
	pos   *Position
	score float64
}

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

func scoreParams1() ScoreParams {
	var weights []int
	for i := 0; i < 60; i++ {
		w := 10.3 - (2.5 * math.Log1p(float64(59-i)))
		// fmt.Printf("%f\n", w)
		weights = append(weights, int(w))
	}
	return ScoreParams{pieceCountWeights: weights}
}
func (s *Scorer) calculateScore() int {
	nrPiecesPlayed := 0
	nrMyPieces := 0
	// count pieces
	for i := 0; i < len(s.b.rows[0]); i++ {
		for j := 0; j < len(s.b.rows); j++ {
			if s.b.rows[i][j] != Empty {
				nrPiecesPlayed += 1
				if s.b.rows[i][j] == s.myPiece {
					nrMyPieces += 1
				}
			}
		}
	}
	return s.params.pieceCountWeights[nrPiecesPlayed]
}
