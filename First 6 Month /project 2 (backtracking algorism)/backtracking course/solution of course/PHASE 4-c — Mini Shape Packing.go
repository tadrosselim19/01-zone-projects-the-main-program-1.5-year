package main

import "fmt"

type shapes struct {
	shape        [][]rune
	all_rotation int
}

var all_shapes []*shapes

// ---------------- DEFINE SHAPE ----------------
func define_shape(shape [][]rune, ch rune) {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] != ' ' && shape[i][j] != '.' {
				shape[i][j] = ch
			}
		}
	}

	rot := 4

	// square → 1 rotation
	if len(shape) == len(shape[0]) {
		rot = 1
	}

	// line → 2 rotations
	if len(shape) == 1 || len(shape[0]) == 1 {
		rot = 2
	}

	all_shapes = append(all_shapes, &shapes{
		shape:        shape,
		all_rotation: rot,
	})
}

// ---------------- ROTATION ----------------
func rotate90(s [][]rune) [][]rune {
	h := len(s)
	w := len(s[0])

	out := make([][]rune, w)
	for i := 0; i < w; i++ {
		out[i] = make([]rune, h)
		for j := 0; j < h; j++ {
			out[i][j] = s[h-1-j][i]
		}
	}
	return out
}

// ---------------- BOARD HELPERS ----------------
func canPlace(s *shapes, bord [][]rune, x, y int) bool {
	for i := 0; i < len(s.shape); i++ {
		for j := 0; j < len(s.shape[i]); j++ {

			if s.shape[i][j] == ' ' {
				continue
			}

			if x+i >= len(bord) || y+j >= len(bord[0]) {
				return false
			}

			if bord[x+i][y+j] != '.' {
				return false
			}
		}
	}
	return true
}

func place(s *shapes, bord [][]rune, x, y int) {
	for i := 0; i < len(s.shape); i++ {
		for j := 0; j < len(s.shape[i]); j++ {
			if s.shape[i][j] != ' ' {
				bord[x+i][y+j] = s.shape[i][j]
			}
		}
	}
}

func remove(s *shapes, bord [][]rune, x, y int) {
	for i := 0; i < len(s.shape); i++ {
		for j := 0; j < len(s.shape[i]); j++ {
			if s.shape[i][j] != ' ' {
				bord[x+i][y+j] = '.'
			}
		}
	}
}

func print_bord(bord [][]rune) {
	for _, r := range bord {
		fmt.Println(string(r))
	}
	fmt.Println()
}

// ---------------- BACKTRACKING ----------------
func mini_shape_packing(index int, bord [][]rune) bool {
	if index == len(all_shapes) {
		print_bord(bord)
		return true
	}

	s := all_shapes[index]
	original := s.shape

	for r := 0; r < s.all_rotation; r++ {
		for i := 0; i < len(bord); i++ {
			for j := 0; j < len(bord[0]); j++ {

				if canPlace(s, bord, i, j) {
					place(s, bord, i, j)

					if mini_shape_packing(index+1, bord) {
						return true
					}

					remove(s, bord, i, j)
				}
			}
		}
		s.shape = rotate90(s.shape)
	}

	s.shape = original
	return false
}

// ---------------- BOARD CREATION ----------------
func make_board(size int) [][]rune {
	b := make([][]rune, size)
	for i := 0; i < size; i++ {
		b[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			b[i][j] = '.'
		}
	}
	return b
}

// ---------------- MAIN ----------------
func main() {
	raw := [][][]rune{
		{
			{'#', '#', '#'},
			{' ', '#', ' '},
		},
		{
			{'#', '#'},
			{'#', '#'},
		},
		{
			{'#', '#', '#'},
		},
	}

	for i, s := range raw {
		define_shape(s, rune('A'+i))
	}

	bord := make_board(6)

	if !mini_shape_packing(0, bord) {
		fmt.Println("NO SOLUTION")
	}
}
