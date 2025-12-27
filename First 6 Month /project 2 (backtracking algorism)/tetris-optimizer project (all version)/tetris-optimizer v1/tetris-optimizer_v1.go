package main

import (
	"fmt"
	"os"
)

type shapes struct {
	shape   []string
	rotation int
}

var all_shapes []*shapes

func split_to_define_shape(s string) {
	s_len := len(s)
	index := 0

	line := ""
	shape := []string{}

	for i := 0; i < s_len; i++ {
		char := string(index + 'A')

		// Skip carriage return (\r)
		if s[i] == '\r' {
			continue
		}

		// Detect blank line (empty line) = end of shape
		if s[i] == '\n' {
			if line == "" {
				// Blank line -> finish current shape
				if len(shape) > 0 {
					all_shapes = append(all_shapes, &shapes{
						shape:   shape,
						rotation: 0,
					})
					index++
					shape = []string{}
				}
				continue
			} else {
				// End of a normal line
				shape = append(shape, line)
				line = ""
				continue
			}
		}

		// Build the current line
		if s[i] == '#' {
			line += char
		} else if s[i] == '.' {
			line += "."
		}
	}

	// Append the last line and shape if file doesn't end with a blank line
	if line != "" {
		shape = append(shape, line)
	}
	if len(shape) > 0 {
		all_shapes = append(all_shapes, &shapes{
			shape:   shape,
			rotation: 0,
		})
	}
}

func can_place(shape []string, bord []string, x, y, N int) bool {
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '.'{
				continue
			}
			if x+i >= N || y+j >= N || bord[x+i][y+j] != '.'{
				return false
			}
		}
	}
	return true
}

func place_shape(shape []string, bord []string, x, y int){
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '.'{
				continue
			}
			// to can chage the byte in string 
			s:= []rune(bord[x+i])
			s[y+j] = rune(shape[i][j])
			bord[x+i] = string(s)
		}
	} 
}

func remove_shape(shape []string, bord []string, x, y int){
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '.'{
				continue
			}
			// to can chage the byte in string 
			s:= []rune(bord[x+i])
			s[y+j] = rune('.')
			bord[x+i] = string(s)
		}
	} 
}
func print_bord(bord []string){
	for _,i:= range(bord){
		fmt.Println(i)
	}
}
func solve_tetris(index int, bord []string,N int) bool {
	if index == len(all_shapes){
		print_bord(bord)
		return true
	}
	
	//for i := 0 ; i < all_shapes[index].rotation ; i++{
		for x := 0 ; x < N ; x++{
			for y := 0 ; y < N ;y++{
				if can_place(all_shapes[index].shape, bord , x,y,N){
					place_shape(all_shapes[index].shape, bord , x,y)
					if solve_tetris(index+1,bord,N){
						return true
					}
					remove_shape(all_shapes[index].shape, bord , x,y)
				}
			}
		}
	//}
	return false
}

func make_bord(n int) []string {
	final := []string{}
	for i := 0; i < n; i++ {
		line := ""
		for j := 0; j < n; j++ {
			line += "."
		}
		final = append(final, line)
	}
	return final
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: usage/ go run. sample.txt")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cant find file")
		return
	}
	file_content := string(file)
	split_to_define_shape(file_content)
	
	stared := 4
	for {
		if solve_tetris(0,make_bord(stared),stared){
			break
		}
		stared++
	}

}
