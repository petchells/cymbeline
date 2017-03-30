package main

type Heuristic struct {
	grid    [][]Square
	myPiece Square
}

func canmove(self Square, opp Square, str [10]Square) bool {
	if str[0] != opp {
		return false
	}
	for ctr := 1; ctr < 8; ctr++ {
		if str[ctr] == Empty {
			return false
		}
		if str[ctr] == self {
			return true
		}
	}
	return false
}

func isLegalMove(self Square, opp Square, grid [][]Square, startx int, starty int) bool {
	if grid[startx][starty] != Empty {
		return false
	}
	str := [10]Square{}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// keep going if both velocities are zero
			if dy == 0 && dx == 0 {
				continue
			}
			//str[0] = '\0'
			for ctr := 1; ctr < 8; ctr++ {
				x := startx + ctr*dx
				y := starty + ctr*dy
				if x >= 0 && y >= 0 && x < 8 && y < 8 {
					str[ctr-1] = grid[x][y]
				} else {
					str[ctr-1] = 0
				}
			}
			if canmove(self, opp, str) {
				return true
			}
		}
	}
	return false
}

func num_valid_moves(self Square, opp Square, grid [][]Square) int {
	count := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if isLegalMove(self, opp, grid, i, j) {
				count += 1
			}
		}
	}
	return count
}

/*
 * Assuming my_color stores your color and opp_color stores opponent's color
 * '-' indicates an empty square on the board
 * 'b' indicates a black tile and 'w' indicates a white tile on the board
 */
func dynamic_heuristic_evaluation_function(grid [][]Square, my_color Square) float64 {
	my_tiles := 0
	opp_tiles := 0
	my_front_tiles := 0
	opp_front_tiles := 0
	opp_color := Black
	if my_color == Black {
		opp_color = White
	}
	// k, my_front_tiles = 0, opp_front_tiles = 0, x, y;

	X1 := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	Y1 := []int{0, 1, 1, 1, 0, -1, -1, -1}

	V := [][]int{
		{20, -3, 11, 8, 8, 11, -3, 20},
		{-3, -7, -4, 1, 1, -4, -7, -3},
		{11, -4, 2, 2, 2, 2, -4, 11},
		{8, 1, 2, -3, -3, 2, 1, 8},
		{8, 1, 2, -3, -3, 2, 1, 8},
		{11, -4, 2, 2, 2, 2, -4, 11},
		{-3, -7, -4, 1, 1, -4, -7, -3},
		{20, -3, 11, 8, 8, 11, -3, 20}}

	// Piece difference, frontier disks and disk squares
	d := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if grid[i][j] == my_color {
				d += V[i][j]
				my_tiles++
			} else if grid[i][j] == opp_color {
				d -= V[i][j]
				opp_tiles += 1
			}
			if grid[i][j] != '-' {
				for k := 0; k < 8; k++ {
					x := i + X1[k]
					y := j + Y1[k]
					if x >= 0 && x < 8 && y >= 0 && y < 8 && grid[x][y] == Empty {
						if grid[i][j] == my_color {
							my_front_tiles += 1
						} else {
							opp_front_tiles += 1
						}
						break
					}
				}
			}
		}
	}
	p := 0.0
	if my_tiles > opp_tiles {
		p = (100.0 * float64(my_tiles)) / float64(my_tiles+opp_tiles)
	} else if my_tiles < opp_tiles {
		p = -(100.0 * float64(opp_tiles)) / float64(my_tiles+opp_tiles)
	} else {
		p = 0
	}

	f := 0.0
	if my_front_tiles > opp_front_tiles {
		f = -(100.0 * float64(my_front_tiles)) / float64(my_front_tiles+opp_front_tiles)
	} else if my_front_tiles < opp_front_tiles {
		f = (100.0 * float64(opp_front_tiles)) / float64(my_front_tiles+opp_front_tiles)
	} else {
		f = 0.0
	}

	// Corner occupancy
	my_tiles = 0
	opp_tiles = 0
	if grid[0][0] == my_color {
		my_tiles += 1
	} else if grid[0][0] == opp_color {
		opp_tiles += 1
	}
	if grid[0][7] == my_color {
		my_tiles += 1
	} else if grid[0][7] == opp_color {
		opp_tiles += 1
	}
	if grid[7][0] == my_color {
		my_tiles += 1
	} else if grid[7][0] == opp_color {
		opp_tiles += 1
	}
	if grid[7][7] == my_color {
		my_tiles += 1
	} else if grid[7][7] == opp_color {
		opp_tiles += 1
	}
	c := 25 * (my_tiles - opp_tiles)

	// Corner closeness
	my_tiles = 0
	opp_tiles = 0
	if grid[0][0] == Empty {
		if grid[0][1] == my_color {
			my_tiles += 1
		} else if grid[0][1] == opp_color {
			opp_tiles += 1
		}
		if grid[1][1] == my_color {
			my_tiles += 1
		} else if grid[1][1] == opp_color {
			opp_tiles += 1
		}
		if grid[1][0] == my_color {
			my_tiles += 1
		} else if grid[1][0] == opp_color {
			opp_tiles += 1
		}
	}
	if grid[0][7] == Empty {
		if grid[0][6] == my_color {
			my_tiles += 1
		} else if grid[0][6] == opp_color {
			opp_tiles += 1
		}
		if grid[1][6] == my_color {
			my_tiles += 1
		} else if grid[1][6] == opp_color {
			opp_tiles += 1
		}
		if grid[1][7] == my_color {
			my_tiles += 1
		} else if grid[1][7] == opp_color {
			opp_tiles += 1
		}
	}
	if grid[7][0] == Empty {
		if grid[7][1] == my_color {
			my_tiles += 1
		} else if grid[7][1] == opp_color {
			opp_tiles += 1
		}
		if grid[6][1] == my_color {
			my_tiles += 1
		} else if grid[6][1] == opp_color {
			opp_tiles += 1
		}
		if grid[6][0] == my_color {
			my_tiles += 1
		} else if grid[6][0] == opp_color {
			opp_tiles += 1
		}
	}
	if grid[7][7] == Empty {
		if grid[6][7] == my_color {
			my_tiles += 1
		} else if grid[6][7] == opp_color {
			opp_tiles += 1
		}
		if grid[6][6] == my_color {
			my_tiles += 1
		} else if grid[6][6] == opp_color {
			opp_tiles += 1
		}
		if grid[7][6] == my_color {
			my_tiles += 1
		} else if grid[7][6] == opp_color {
			opp_tiles += 1
		}
	}
	l := -12.5 * float64(my_tiles-opp_tiles)

	// Mobility
	m := 0.0
	my_tiles = num_valid_moves(my_color, opp_color, grid)
	opp_tiles = num_valid_moves(opp_color, my_color, grid)
	if my_tiles > opp_tiles {
		m = (100.0 * float64(my_tiles)) / float64(my_tiles+opp_tiles)
	} else if my_tiles < opp_tiles {
		m = -(100.0 * float64(opp_tiles)) / float64(my_tiles+opp_tiles)
	} else {
		m = 0
	}

	// final weighted score
	score := (10 * p) + (801.724 * float64(c)) + (382.026 * l) + (78.922 * m) + (74.396 * f) + float64(10*d)
	return score
}
