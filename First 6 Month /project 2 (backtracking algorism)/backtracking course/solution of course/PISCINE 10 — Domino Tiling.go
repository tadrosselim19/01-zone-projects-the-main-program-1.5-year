package main

import (
	"fmt"
	
)

func DOMINO_TILING(bord [][]int, check [][]int, x1 int, y1 int,x2 int, y2 int) bool {
	all_done := true
	x := 0
	y := 0
	for i := 0; i < len(check); i++ {
		for j := 0; j < len(check[i]); j++ {
			if check[i][j] == 0{
				x = i
				y = j
				all_done = false
				break
			}
		}
		if all_done==false{
		break
		}
	}
	if all_done == true{
		return true
	}
	if y+1 < len(bord[0]) &&
		bord[x][y] != -1 &&
		bord[x][y+1] != -1 &&
		check[x][y] == 0 &&
		check[x][y+1] == 0 {

		check[x][y] = 1
		check[x][y+1] = 1

		if DOMINO_TILING(bord, check, x, y, x, y+1) {
			return true
		}

		check[x][y] = 0
		check[x][y+1] = 0
	}

	// Try placing vertically: (x,y) & (x+1,y)
	if x+1 < len(bord) &&
		bord[x][y] != -1 &&
		bord[x+1][y] != -1 &&
		check[x][y] == 0 &&
		check[x+1][y] == 0 {

		check[x][y] = 1
		check[x+1][y] = 1

		if DOMINO_TILING(bord, check, x, y, x+1, y) {
			return true
		}

		check[x][y] = 0
		check[x+1][y] = 0
	}

	return false
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if DOMINO_TILING(grid, grid, 0, 0,0,0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
