package main

import "fmt"

func Sum_K_Constraint(index int,limit_num int, limit_sum int,sum int, result []int) {
	if index > limit_num{
		fmt.Println(result)
		//fmt.Println(index," ",result," ",sum)
		return
	}
	Sum_K_Constraint(index+1,limit_num,limit_sum,sum,result)
	if index + sum <= limit_sum{
		result = append(result, index)
		sum+=index
		//fmt.Println(index," ",result," ",sum)
		Sum_K_Constraint(index+1,limit_num,limit_sum,sum,result)
		sum-=result[len(result)-1]
		result=result[:len(result)-1]
		//fmt.Println(index," ",result," ",sum)
	}
}
func main() {
	limit_num := 4
	limit_sum := 5
	Sum_K_Constraint(1,limit_num, limit_sum,0, []int{})
}
