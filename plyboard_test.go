// test
package main

import (
	"fmt"
	"testing"
)

func TestPlyBoardFull(t *testing.T) {

	board := newBoard()
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board.rows[i][j] = White
		}
	}
	board.rows[0][0] = Black
	board.rows[0][7] = Empty
	board.rows[1][7] = Black
	pb := PlyBoard{dynamic_heuristic_evaluation_function}
	score := dynamic_heuristic_evaluation_function(board.rows, Black)
	move := pb.deepSearch(board, Black)
	if move.score != -12416.295634920636 {
		t.Errorf("Expected %f but got %f", -12416.295634920636, move.score)
	}
	if score != -12416.295634920636 {
		t.Errorf("Expected %f but got %f", -12416.295634920636, score)
	}
}

func TestPlyBoardStart(t *testing.T) {

	board := newBoard()
	pb := PlyBoard{dynamic_heuristic_evaluation_function}
	score := dynamic_heuristic_evaluation_function(board.rows, Black)
	move := pb.deepSearch(board, Black)
	fmt.Println("move", move.pos.AsString(), move.score)
	if move.score != -12416.295634920636 {
		t.Errorf("Expected %f but got %f", -12416.295634920636, move.score)
	}
	if score != -12416.295634920636 {
		t.Errorf("Expected %f but got %f", -12416.295634920636, score)
	}
}

func BenchmarkPlyBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
