package main

import (
	//"log"
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

func (pb *PlyBoard) alphabeta(b *Board, myColour Square, depth int) Move {

	var recurse func(*Board, Square, int, float64, float64) float64

	recurse = func(b1 *Board, colour Square, d int, alpha float64, beta float64) float64 {
		if d == 0 {
			//log.Println("evaluating for ", enumToColour(colour))
			return pb.evaluationFunction(b1.rows, colour)
		}
		bcp := &Board{}
		oppColour := pb.oppColour(colour)
		if colour == myColour {
			vpos := b.findAllValidMoves(colour)
			v := math.Inf(-1)
			for _, mv := range vpos {
				bcp.copyFrom(b1)
				bcp.playMove(&mv, colour)
				v := math.Max(v, recurse(bcp, oppColour, d-1, alpha, beta))
				alpha = math.Max(alpha, v)
				if beta <= alpha {
					return v // beta cut-off
				}
			}
			return v
		} else {
			vpos := b.findAllValidMoves(oppColour)
			v := math.Inf(+1)
			for _, mv := range vpos {
				bcp.copyFrom(b1)
				bcp.playMove(&mv, oppColour)
				v := math.Min(v, recurse(bcp, colour, d-1, alpha, beta))
				beta = math.Min(alpha, v)
				if beta <= alpha {
					return v // alpha cut-off
				}
			}
			return v
		}
	}

	moves := pb.findAllMoves(b, myColour)
	if len(moves) == 0 {
		return Move{score: 0.0, pos: &Position{-1, -1}}
	} else if len(moves) == 1 {
		return moves[0]
	}
	// first 'alpha' is sorted -- should make pruning more efficient
	sort.Sort(ByScore(moves))
	beta := math.Inf(+1)
	bcp := &Board{}
	oppColour := pb.oppColour(myColour)
	//log.Println("================")
	for _, mv := range moves {
		//log.Println("Move: ", mv.pos.AsString(), mv.score)
		bcp.copyFrom(b)
		bcp.playMove(mv.pos, myColour)
		mv.score = recurse(bcp, oppColour, depth-1, math.Inf(-1), beta)
		//log.Println("Move: ", mv.pos.AsString(), mv.score)

		//		alpha = math.Max(alpha, v)
		//		if beta <= alpha {
		//			return mv // beta cut-off
		//		}
	}
	sort.Sort(ByScore(moves))
	return moves[0]

}

func (pb *PlyBoard) deepSearch(b *Board, myColour Square, depth int) Move {
	return pb.alphabeta(b, myColour, 3)
}

func (pb *PlyBoard) findAllMoves(b *Board, myColour Square) []Move {
	moves := []Move{}
	positions := b.findAllValidMoves(myColour)
	bcp := &Board{}
	for _, pos := range positions {
		bcp.copyFrom(b)
		bcp.playMove(&pos, myColour)
		score := pb.evaluationFunction(bcp.rows, myColour)
		moves = append(moves, Move{pos: &Position{pos.x, pos.y}, score: score})
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
