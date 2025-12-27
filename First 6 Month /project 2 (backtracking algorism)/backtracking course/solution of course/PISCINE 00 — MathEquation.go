package main

import "fmt"


var memo map[[2]int][]string
func Math_Equation_check(arr []int, i int, result int, sum int) bool {
	if i == len(arr) {
		return sum == result
	}
	return Math_Equation_check(arr, i+1, result+arr[i], sum) || Math_Equation_check(arr, i+1, result-arr[i], sum)
}

// backtracking tech
func Math_Equation_solve(arr []int, i int, result int, sum int, final []string) []string {
	if i == len(arr) {
		if result == sum {
			return append([]string{}, final...) // return copy
		}
		return nil
	}

	// Try "+"
	final = append(final, "+")
	ans := Math_Equation_solve(arr, i+1, result+arr[i], sum, final)
	if ans != nil {
		return ans
	}
	final = final[:len(final)-1] // backtrack

	// Try "-"
	final = append(final, "-")
	ans = Math_Equation_solve(arr, i+1, result-arr[i], sum, final)
	if ans != nil {
		return ans
	}
	final = final[:len(final)-1] // backtrack

	return nil
}


// recursion only
func Math_Equation_solver(arr []int, i int, result int, sum int, final *[5]string) *[5]string {
	if i == len(arr) {
		if sum == result {
			return final
		}
		return nil
	}
	final[i] = "+"
	ans := Math_Equation_solver(arr, i+1, result+arr[i], sum, final)
	if ans != nil {
		return ans
	}
	final[i] = "-"
	ans = Math_Equation_solver(arr, i+1, result-arr[i], sum, final)
	if ans != nil {
		return ans
	}
	return nil
}


// dynamic programing
func Math_DP(arr []int, i int, result int, target int) []string {

	key := [2]int{i, result}
	if v, ok := memo[key]; ok {
		return v // cached result (can be nil or solution)
	}

	if i == len(arr) {
		if result == target {
			return []string{}
		}
		memo[key] = nil
		return nil
	}

	// Try "+"
	ans := Math_DP(arr, i+1, result+arr[i], target)
	if ans != nil {
		memo[key] = append([]string{ "+" }, ans...)
		return memo[key]
	}

	// Try "-"
	ans = Math_DP(arr, i+1, result-arr[i], target)
	if ans != nil {
		memo[key] = append([]string{ "-" }, ans...)
		return memo[key]
	}
	memo[key] = nil
	return nil
}



func main() {

	arr := []int{7, 3, 5, 4, 2}
	sum := -21
	opp := [5]string{}
	fmt.Println(Math_Equation_check(arr, 0, 0, sum))
	fmt.Println(Math_Equation_solve(arr, 0, 0, sum, []string{}))
	fmt.Println(*Math_Equation_solver(arr, 0, 0, sum, &opp))

	memo = make(map[[2]int][]string)

	result := Math_DP(arr, 0, 0, sum)

	fmt.Println("DP Solution:", result)
	

}
