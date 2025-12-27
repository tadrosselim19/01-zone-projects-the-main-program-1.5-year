package main

import "fmt"

type shapes struct {
	shape        [][]rune
	all_rotation int
}

var all_shapes []*shapes

func define_shape(shape [][]rune, ch rune) {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] != ' ' && shape[i][j] != '.' {
				shape[i][j] = ch
			}
		}
	}
	// square shape
	if len(shape) == len(shape[0]) {
		sh := &shapes{
			shape:        shape,
			all_rotation: 1,
		}
		all_shapes = append(all_shapes, sh)
		return
	}
	// line shape
	if len(shape) == 1 {
		sh := &shapes{
			shape:        shape,
			all_rotation: 2,
		}
		all_shapes = append(all_shapes, sh)
		return
	}
	for _, i := range shape {
		if len(i) != 1 {
			sh := &shapes{
				shape:        shape,
				all_rotation: 4,
			}
			all_shapes = append(all_shapes, sh)
			return
		}
	}
	sh := &shapes{
		shape:        shape,
		all_rotation: 2,
	}
	all_shapes = append(all_shapes, sh)
	return

}

// main tracking
func find_ans(bord [][]rune) bool {
	count := 0
	for _, i := range all_shapes {
		for _, j := range bord {
			find := false
			for _, k := range j {
				if k == i.shape[0][0] {
					find = true
					count++
					break
				}
			}
			if find {
				break
			}

		}
	}
	if count == len(all_shapes) {
		return true
	}
	return false
}

func canPlace(s *shapes, bord [][]rune, x, y int) bool {
	for i := 0; i < len(s.shape); i++ {
		for j := 0; j < len(s.shape[i]); j++ {
			if s.shape[i][j] == ' ' {
				continue
			}
			if x+i >= len(bord) || y+j >= len(bord[i]) {
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

func main() {
	// shapes := [][][]rune{
	// 	{{'#', '#', '#'},
	// 		{' ', '#', ' '}},

	// 	{{'#', '#'},
	// 		{'#', '#'}},

	// 	{{'#', '#', '#'}},
	// }

	shapes := [][][]rune{
		{{'#'},
			{'#'},
			{'#'},
			{'#'}},

		{{'#', '#', '#'}},

		{{'#', '#', '#'},
			{' ', ' ', '#'}},

		{{' ', '#', '#'},
			{'#', '#', ' '}},

		{{'#', '#'},
			{'#', '#'}},

		{{'#', '#', ' '},
			{' ', '#', '#'}},

		{{'#', '#'},
			{' ', '#'},
			{' ', '#'},},

			
		{{'#', '#', '#'},
			{' ', '#', ' '}},

	}

	// size of bord
	bord_size := 0

	for index, shape := range shapes {
		define_shape(shape, rune(index+'A'))
		if len(shape) >= len(shape[0]) {
			bord_size += len(shape)
		} else {
			bord_size += len(shape[0])
		}
	}
	bord_size = bord_size / 4
	bord := make_board(bord_size)

	if !mini_shape_packing(0, bord) {
		fmt.Println("NO SOLUTION")
	}

}

// ROOT: func main()
// │
// ├── PHASE 1: PREPARATION (Data Setup)
// │   │
// │   ├── Step 1: Define Raw Data
// │   │   └── Create a slice of raw '#' shapes (Tetris-like blocks)
// │   │
// │   ├── Step 2: Initialize Shapes Loop
// │   │   └── For each shape in the list:
// │   │       └── Call func define_shape()
// │   │           ├── Convert '#' to a unique Letter (A, B, C...)
// │   │           └── Optimization: Calculate 'all_rotation'
// │   │               ├── Is it a square? -> Rotate 1 time (Symmetry)
// │   │               ├── Is it a line?   -> Rotate 2 times
// │   │               └── Irregular?      -> Rotate 4 times
// │   │
// │   └── Step 3: Initialize Board
// │       ├── Calculate required board size
// │       └── Call func make_board()
// │           └── Create a 2D grid filled with dots ('.')
// │
// ├── PHASE 2: THE SOLVER ENGINE (Recursive Backtracking)
// │   │
// │   └── Call func mini_shape_packing(index=0, board)
// │       │
// │       ├── [BASE CASE: SUCCESS]
// │       │   ├── Condition: Have we placed all shapes? (index == len)
// │       │   ├── Action: Call print_bord()
// │       │   └── Return TRUE (Stop everything, we found the answer!)
// │       │
// │       └── [RECURSIVE STEP: TRYING TO FIT]
// │           │
// │           ├── Loop 1: Rotation Attempts (0 to s.all_rotation)
// │           │   │
// │           │   ├── Loop 2: Board Rows (x)
// │           │   │   │
// │           │   │   └── Loop 3: Board Columns (y)
// │           │   │       │
// │           │   │       ├── Check: func canPlace(shape, x, y)
// │           │   │       │   ├── Is it inside bounds?
// │           │   │       │   └── Is the spot empty ('.')?
// │           │   │       │
// │           │   │       └── IF YES (Valid Spot):
// │           │   │           ├── Action: func place() (Write letter to board)
// │           │   │           │
// │           │   │           ├── RECURSE: mini_shape_packing(index + 1)
// │           │   │           │   └── If this returns TRUE -> Return TRUE
// │           │   │           │
// │           │   │           └── BACKTRACK: func remove() (Write '.' back to board)
// │           │   │               (We only reach here if the recursion failed)
// │           │   │
// │           │   └── Action: func rotate90()
// │           │       └── Rotate the shape matrix to try a new angle
// │           │
// │           └── [FAILURE]
// │               └── If no position/rotation works, return FALSE
// │
// └── PHASE 3: OUTPUT
//     └── If mini_shape_packing returns false
//         └── Print "NO SOLUTION"