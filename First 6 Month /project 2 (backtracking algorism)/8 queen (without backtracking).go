package main

import (
	"fmt"
)

const size = 7

func main() {
	bord := make_bord()
	solve(bord)

}
func make_bord() [][]rune {
	bord := [][]rune{{'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}, {'-', '-', '-', '-', '-', '-', '-', '-'}}
	return bord
}

func print_solved_bord(bord [][]rune) {
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			fmt.Print(string(bord[i][j]) + " ")
		}
		println()
	}
}

func solve(bord [][]rune) {
	

	for k := 0; k <= size; k++ {
		for l := 0; l <= size; l++ {
			bord = make_bord()
			bord[k][l] = 'q'

			for i := 0; i <= size; i++ {
				for j := 0; j <= size; j++ {
					if is_correct(bord, i, j) == true {
						bord[i][j] = 'q'
					}
				}
			}
			if is_finshed(bord) == true {
				print_solved_bord(bord)

			}
		}
	}

}

func is_correct(bord [][]rune, row int, collum int) bool {
	if bord[row][collum] == '-' {
		for i := 0; i <= size; i++ {
			if bord[row][i] == 'q' {
				return false
			}
		}
		for i := 0; i <= size; i++ {
			if bord[i][collum] == 'q' {
				return false
			}
		}

		// Check upper-left diagonal ↖
		for i, j := row-1, collum-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if bord[i][j] == 'q' {
				return false
			}
		}

		// Check upper-right diagonal ↗
		for i, j := row-1, collum+1; i >= 0 && j < len(bord); i, j = i-1, j+1 {
			if bord[i][j] == 'q' {
				return false
			}
		}

		// Check lower-left diagonal ↙
		for i, j := row+1, collum-1; i < len(bord) && j >= 0; i, j = i+1, j-1 {
			if bord[i][j] == 'q' {
				return false
			}
		}

		// Check lower-right diagonal ↘
		for i, j := row+1, collum+1; i < len(bord) && j < len(bord); i, j = i+1, j+1 {
			if bord[i][j] == 'q' {
				return false
			}
		}

	}
	return true

}

func is_finshed(bord [][]rune) bool {
	count_q := 0
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			if bord[i][j] == 'q' {
				count_q++
			}
		}
	}

	return count_q == 8
}
