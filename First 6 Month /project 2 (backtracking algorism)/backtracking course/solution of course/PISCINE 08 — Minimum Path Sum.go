package main

import (
	"fmt"
	"strconv"
)

var results = make(map[int]string)

func Path_Sum(grid [][]int, x int, y int, end []int, path string, sum int) {
	// FIX 1: correct bounds
	if x > end[0] || y > end[1] {
		return
	}

	path = path + "(" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ") "
	sum += grid[x][y]

	// FIX 2: correct end check
	if x == end[0] && y == end[1] {
		results[sum] = path
		return
	}

	Path_Sum(grid, x+1, y, end, path, sum)
	Path_Sum(grid, x, y+1, end, path, sum)
	Path_Sum(grid, x+1, y+1, end, path, sum)
}

func Minimum_Path_Sum(grid [][]int, x int, y int, end []int, path string, sum int) {
	Path_Sum(grid, x, y, end, "", 0)

	var min int
	started := false

	for key := range results {
		if !started {
			min = key
			started = true
		}
		if key < min {
			min = key
		}
	}

	fmt.Println("Path:", results[min])
	fmt.Println("Min Sum:", min)
}

func main() {
	input := [][]int{
		{1, 3},
		{2, 1},
		{4, 2},
	}
	start := []int{0, 0}
	end := []int{2, 1} // FIXED

	Minimum_Path_Sum(input, start[0], start[1], end, "", 0)
}
