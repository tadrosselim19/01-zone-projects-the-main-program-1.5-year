package main

import (
	"fmt"
)

func is_correct(N int, bord [][]string, x int, y int) bool {
	// Check row
	for j := 0; j < N; j++ {
		if bord[x][j] == "q" {
			return false
		}
	}

	// Check column
	for i := 0; i < N; i++ {
		if bord[i][y] == "q" {
			return false
		}
	}

	// UL diagonal
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if bord[i][j] == "q" {
			return false
		}
	}

	// UR diagonal
	for i, j := x-1, y+1; i >= 0 && j < N; i, j = i-1, j+1 {
		if bord[i][j] == "q" {
			return false
		}
	}

	// LL diagonal
	for i, j := x+1, y-1; i < N && j >= 0; i, j = i+1, j-1 {
		if bord[i][j] == "q" {
			return false
		}
	}

	// LR diagonal
	for i, j := x+1, y+1; i < N && j < N; i, j = i+1, j+1 {
		if bord[i][j] == "q" {
			return false
		}
	}
	return true
}

func N_Queens(N int, bord [][]string, row int, queen_count int) {
    if queen_count == N { // solution found
        for _, r := range bord {
            fmt.Println(r)
        }
        fmt.Println()
        return
    }
    if row >= N { // past last row
        return
    }

    for col := 0; col < N; col++ {
        if is_correct(N, bord, row, col) {
            bord[row][col] = "q"
            N_Queens(N, bord, row+1, queen_count+1) // next row
            bord[row][col] = "_" // backtrack
        }
    }
}


func main() {
	const N = 8
	bord := [][]string{}
	for i := 0; i < N; i++ {
		bord = append(bord, []string{"_", "_", "_", "_","_", "_", "_", "_",})
	}

	N_Queens(N, bord, 0, 0)
}
// (row 0)
// ├─ (0,0)
// │  ├─ (1,2)
// │  │  ├─ (2,1)
// │  │  │  └─ (3,3) → solution 1
// │  │  └─ (2,3) X (blocked)
// │  └─ (1,3)
// │     ├─ (2,1) X (blocked)
// │     └─ (2,2)
// │        └─ (3,0) → solution 2
// ├─ (0,1)
// │  └─ ...
// ├─ (0,2)
// │  └─ ...
// └─ (0,3)
//    └─ ...
