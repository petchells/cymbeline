package main

import (
	//"log"
	//"fmt"
	"math"
	"math/rand"
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
	//log.Println("alphabeta", enumToColour(myColour))
	recurse = func(b1 *Board, colour Square, d int, alpha float64, beta float64) float64 {
		//log.Println("evaluating for ", enumToColour(colour), "at depth", d)
		oppColour := pb.oppColour(colour)
		if d == 0 {
			return pb.evaluationFunction(b1.rows, myColour)
		}
		bcp := &Board{}
		if colour == myColour {
			v := math.Inf(-1)
			vpos := b1.findAllValidMoves(colour)
			if len(vpos) == 0 {
				//				vposOpp := b1.findAllValidMoves(oppColour)
				//				if len(vposOpp) == 0 {
				return pb.evaluationFunction(b1.rows, myColour)
				//				} else {
				//					v = math.Max(v, recurse(bcp, oppColour, d, alpha, beta))
				//				}
			} else {
				for _, mv := range vpos {
					bcp.copyFrom(b1)
					bcp.playMove(&mv, colour)
					v = math.Max(v, recurse(bcp, oppColour, d-1, alpha, beta))
					alpha = math.Max(alpha, v)
					if beta <= alpha {
						// log.Println("beta cut-off")
						return v // beta cut-off
					}
				}
			}
			return v
		} else {
			v := math.Inf(+1)
			vpos := b.findAllValidMoves(colour)
			if len(vpos) == 0 {
				//				vposOpp := b1.findAllValidMoves(oppColour)
				//				if len(vposOpp) == 0 {
				return pb.evaluationFunction(b1.rows, oppColour)
				//				} else {
				//					v = math.Max(v, recurse(bcp, oppColour, d, alpha, beta))
				//				}
			} else {
				for _, mv := range vpos {
					bcp.copyFrom(b1)
					bcp.playMove(&mv, colour)
					v = math.Min(v, recurse(bcp, myColour, d-1, alpha, beta))
					beta = math.Min(alpha, v)
					if beta <= alpha {
						// log.Println("alpha cut-off")
						return v // alpha cut-off
					}
				}
			}
			return v
		}
	}

	moves := pb.findAllMoves(b, myColour)
	if len(moves) == 0 {
		return Move{score: 0.0}
	} else if len(moves) == 1 {
		return moves[0]
	}
	// first 'alpha' is sorted -- should make pruning more efficient
	sort.Sort(ByScore(moves))
	alpha, beta := math.Inf(-1), math.Inf(+1)
	bcp := &Board{}
	oppColour := pb.oppColour(myColour)
	for _, mv := range moves {
		bcp.copyFrom(b)
		bcp.playMove(mv.pos, myColour)
		mv.score = recurse(bcp, oppColour, depth-1, alpha, beta)
	}
	sort.Sort(ByScore(moves))
	//fmt.Println("---")
	var best []Move
	for i, mv := range moves {
		//fmt.Println("mv", mv.pos.AsString(), mv.score)
		if i == 0 {
			best = append(best, mv)
		} else {
			if mv.score >= best[0].score {
				best = append(best, mv)
			}
		}
	}
	//	for _, mv := range best {
	//		fmt.Println("best", mv.pos.AsString(), mv.score)
	//	}
	return best[rand.Intn(len(best))]
}

func (pb *PlyBoard) deepSearch(b *Board, myColour Square) Move {
	return pb.alphabeta(b, myColour, 3)
}

/**
 * Find the best move for a given board position. Single ply
 */
func (pb *PlyBoard) findBestMove(b *Board, myPiece Square) Move {
	moves := pb.findAllMoves(b, myPiece)
	if len(moves) == 0 {
		return Move{}
	}
	if len(moves) == 1 {
		return moves[0]
	}
	sort.Sort(ByScore(moves))
	//	fmt.Println("---")
	var best []Move
	for i, mv := range moves {
		//		fmt.Println("mv", mv.pos.AsString(), mv.score)
		if i == 0 {
			best = append(best, mv)
		} else {
			if mv.score >= best[0].score {
				best = append(best, mv)
			}
		}
	}
	//	for _, mv := range best {
	//		fmt.Println("best", mv.pos.AsString(), mv.score)
	//	}
	return best[rand.Intn(len(best))]
}
func (pb *PlyBoard) findBestMoveDither(b *Board, myPiece Square) Move {
	// Iterate over the top level moves
	var move Move
	bestScore := -99999999.0
	validMoves := b.findAllValidMoves(myPiece)
	nrBlack, nrWhite := b.countPieces()
	randWeight := 1e-14 / float64(nrBlack+nrWhite)
	for i := 0; i < len(validMoves); i++ {
		bcp := b.copy()
		bcp.playMove(&validMoves[i], myPiece)
		score := pb.evaluationFunction(bcp.rows, myPiece)
		// randomize the score a bit
		rand := (rand.Float64() - .5) * randWeight * (10. + score)
		if score+rand > bestScore {
			bestScore = score
			move = Move{pos: &validMoves[i], score: score}
		}
	}
	if bestScore == -99999999.0 {
		return Move{}
	}
	return move
}
func (pb *PlyBoard) findBestMoveRandom(b *Board, myColour Square) Move {
	// Iterate over the top level moves
	validMoves := []Move{}
	totalScores := 0.0
	for i := 0; i < len(b.rows[0]); i++ {
		for j := 0; j < len(b.rows); j++ {
			pos := &Position{
				x: i,
				y: j}
			if b.isValidMove(pos, myColour) {
				bcp := b.copy()
				bcp.playMove(pos, myColour)
				score := dynamic_heuristic_evaluation_function_alt(bcp.rows, myColour)
				totalScores += score
				validMoves = append(validMoves, Move{pos: pos, score: score})
			}
		}
	}
	if len(validMoves) == 0 {
		return Move{}
	}
	sort.Sort(ByScore(validMoves))
	pick := rand.Float64() * totalScores
	for i := 0; i < len(validMoves); i++ {
		totalScores -= validMoves[i].score
		if pick >= totalScores {
			return validMoves[i]
		}
	}
	var move Move
	if l := len(validMoves) / 2; l == 0 {
		move = validMoves[0]
	} else {
		move = validMoves[rand.Intn(l)]
	}
	return move
}

func (pb *PlyBoard) findAllMoves(b *Board, myColour Square) []Move {
	moves := []Move{}
	positions := b.findAllValidMoves(myColour)
	bcp := &Board{}
	for _, pos := range positions {
		bcp.copyFrom(b)
		bcp.playMove(&pos, myColour)
		move := Move{pos: &Position{pos.x, pos.y}, score: pb.evaluationFunction(bcp.rows, myColour)}
		// log.Println("move", move.pos.AsString(), move.score)
		// bcp.printboard()
		moves = append(moves, move)
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
