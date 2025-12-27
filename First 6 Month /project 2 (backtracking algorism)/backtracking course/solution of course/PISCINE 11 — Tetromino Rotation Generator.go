package main

import (
	"fmt"
)

func rotate(shape [][]rune, degree int) [][]rune {
	final := [][]rune{}
	if degree == 90 {
		for i := 0; i < len(shape[0]); i++ {
			temp := []rune{}
			for j := len(shape) - 1; j >= 0; j-- {
				temp = append(temp, shape[j][i])
			}
			final = append(final, temp)
		}
	}
	if degree == 180 {
		for i := len(shape) - 1; i >= 0; i-- {
			temp := []rune{}
			for j := len(shape[i]) - 1; j >= 0; j-- {
				temp = append(temp, shape[i][j])
			}
			final = append(final, temp)
		}
	}

	if degree == 270 {
		for i := len(shape[0]) - 1; i >= 0; i-- {
			temp := []rune{}
			for j := 0; j < len(shape); j++ {
				temp = append(temp, shape[j][i])
			}
			final = append(final, temp)
		}
	}
	return final
}

func toString(shape [][]rune) string {
	final := ""
	rows := len(shape)
	cols := len(shape[0])

	// find bounding box max dimension
	box := rows
	if cols > box {
		box = cols
	}

	// print in a square box
	for i := 0; i < box; i++ {
		for j := 0; j < box; j++ {

			if i < rows && j < cols {
				if shape[i][j] != ' ' {
					final += string(shape[i][j])
				} else {
					final += "."
				}
			} else {
				final += "."
			}
		}
		final += "\n"
	}
	return final
}

func generateRotations(shape [][]rune) {
	seen := make(map[string]bool)

	// rotation 0
	fmt.Printf("--- rotation %d ---\n", 0)
	s0 := toString(shape)
	fmt.Print(s0)
	seen[s0] = true

	// 90, 180, 270
	rotations := []int{90, 180, 270}
	for _, deg := range rotations {
		r := rotate(shape, deg)
		str := toString(r)

		if !seen[str] {
			fmt.Printf("--- rotation %d ---\n", deg)
			fmt.Print(str)
			seen[str] = true
		}
	}
}

func main() {
	shape := [][]rune{
		{'o', '#'},
		{'#', '#'},
		{'#', '#'},
		{'#', '#'},
	}

	shape1 := [][]rune{
		{'#', ' '},
		{'#', ' '},
		{'#', 'o'},
		{'#', ' '},
	}

	shape2 := [][]rune{
		{'#', '#','#'},
		{' ', '#',' '},
		{' ', '#',' '},

		
	}

	generateRotations(shape)
	println("--------------------")
	generateRotations(shape1)
	println("--------------------")
	generateRotations(shape2)
}
