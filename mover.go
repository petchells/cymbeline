package main

/**
 * Find the best move for a given board position
 */

func findBestMove(b *Board, colour Square) {
	// Iterate over the top level moves

}

func isValidMove(b *Board, p *Position, mySide Square) bool {
	if b.getSquare(p) != Empty {
		return false
	}
	// must have an opposing piece next to it

	return true
}
