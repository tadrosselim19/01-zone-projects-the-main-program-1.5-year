package main

import (
	"fmt"
	"strconv"
)

func Knights_Tour(bord [][]string, x int, y int,tour [][]int,path string,count_tour int) {
	if x >= len(bord) || y >= len(bord) ||x < 0 || y < 0 || tour[x][y] != -1{
		return
	}
	tour[x][y]=1
	path = path + "(" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ")\n"
	correct := true
	for i := 0 ; i < len(tour) ; i++{
		for j := 0 ; j < len(tour[i]);j++{
			if tour[i][j]==-1{
				correct = false
				break
			}

		}
	}
	if correct == true && count_tour == len(bord)* len(bord) -1 {
		fmt.Println(path)
		tour[x][y] = -1 //if you want the frist succes tour remove this line 
		return
	}
	
	Knights_Tour(bord,x+2,y+1,tour,path,count_tour+1)
	
	Knights_Tour(bord,x+2,y-1,tour,path,count_tour+1)
	
	Knights_Tour(bord,x-2,y+1,tour,path,count_tour+1)
	
	Knights_Tour(bord,x-2,y-1,tour,path,count_tour+1)
	
	Knights_Tour(bord,x+1,y+2,tour,path,count_tour+1)
	
	Knights_Tour(bord,x+1,y-2,tour,path,count_tour+1)
	
	Knights_Tour(bord,x-1,y+2,tour,path,count_tour+1)
	
	Knights_Tour(bord,x-1,y-2,tour,path,count_tour+1)
	tour[x][y]=-1

}
func main() {
	input := [][]string{
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
	}
	
	start := []int{0, 0}
	tour := [][]int{}
	
	for i := 0 ; i < len(input) ; i++{
		tour = append(tour,[]int{-1,-1,-1,-1,-1} )
	}
	
	Knights_Tour(input,start[0],start[1],tour,"",0)
	fmt.Println(tour)
}