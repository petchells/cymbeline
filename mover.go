package main

/**
 * Find the best move for a given board position
 */

func findBestMove(b *Board, colour Square) {
	// Iterate over the top level moves

}
func isOnBoard(b *Board, p *Position) bool {
	return p.x >= 0 && p.y >= 0 &&
		p.x < int8(len(b.rows)) && p.y < int8(len(b.rows))
}
func scanDiagonal(b *Board, p *Position, myPiece Square, xinc int8, yinc int8) []Position {
	turned := make([]Position, 0, len(b.rows))
	foundOpp := false
	for true {
		nextPos := Position{x: p.x + xinc, y: p.x + yinc}
		if isOnBoard(b, &nextPos) {
			// nothing to see here
			continue
		}
		if b.getSquare(&nextPos) == Empty {
			// none of our guys in that line
			continue
		}
		if b.getSquare(&nextPos) != myPiece {
			turned = append(turned, nextPos)
			foundOpp = foundOpp || true
			// otherwise, this is another in a line of opponents.
		} else {
			if foundOpp {
				// we've come to the end of the oppenents
				return turned
			} else {
				// nothing turned
				continue
			}
		}
	}
	return []Position{}
}
func scanAllDiagonals(b *Board, p *Position, myPiece Square) []Position {
	turned := make([]Position, 0, len(b.rows)*4)
	for i := int8(-1); i <= 1; i++ {
		for j := int8(-1); j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			positions := scanDiagonal(b, p, myPiece, i, j)
			if len(positions) > 0 {
				turned = append(turned, positions...)
			}
		}
	}
	return turned
}
func playMove(b *Board, p *Position, myPiece Square) bool {
	if b.getSquare(p) != Empty {
		return false
	}
	// look along all diagonals for turned pieces
	positions := scanAllDiagonals(b, p, myPiece)
	if len(positions) == 0 {
		// no pieces were actually turned
		return false
	}
	b.setPiece(p, myPiece)
	for _, pos := range positions {
		b.setPiece(&pos, myPiece)
	}
	return true
}
func isValidMove(b *Board, p *Position, myPiece Square) bool {
	if b.getSquare(p) != Empty {
		return false
	}
	// must have an opposing piece next to it

	return true
}
