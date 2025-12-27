package main

import "fmt"

func Obstacle_Grid_Path_Finder(bord [][]string, start []int, end []int, x int, y int,obstacle [][]int, path string) {
	if x > end[0] || y > end[1] {
		return
	}
	for i := 0 ; i < len(obstacle); i++{
		if obstacle[i][0] == y && obstacle[i][1]== x{
			return
		}
	}
	if x == end[0] && y == end[1] {
		fmt.Println(path)
		return
	}
	Obstacle_Grid_Path_Finder(bord,start,end,x+1,y,obstacle,path+"R")
	
	
	Obstacle_Grid_Path_Finder(bord,start,end,x,y+1,obstacle,path + "D")
	
}
func main() {
	input := [][]string{
		{".", ".", "."},
		{".", "T", "."},
		{".", ".", "."},
	}
	
	start := []int{0, 0}
	end := []int{2, 2}
	obstacle := [][]int{
		{1,1},
	}
	obstacle2 := [][]int{
		{1,0},
		{1,2},
	}
	Obstacle_Grid_Path_Finder(input, start, end, start[0], start[1],obstacle, "")
	println()
	Obstacle_Grid_Path_Finder(input, start, end, start[0], start[1],obstacle2, "")

}