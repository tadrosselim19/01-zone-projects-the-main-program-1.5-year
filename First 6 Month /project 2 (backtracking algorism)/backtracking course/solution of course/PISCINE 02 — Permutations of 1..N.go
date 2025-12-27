package main

import "fmt"

func backtrack(n int, perm []int, used []bool) {

    // BASE CASE
    if len(perm) == n{
		fmt.Println(perm)
		return
	}
    // TRY ALL NUMBERS 1..n
    for i := 1; i <= n; i++ {
		if used[i] == false{
			perm = append(perm, i)
			used[i] = true
		}else{
			continue
		}
        backtrack(n, perm, used)
		used[i] = false
		perm =perm[:len(perm)-1]
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    // Start with empty permutation
    perm := []int{}

    // All numbers unused initially
    used := make([]bool, n+1)

    backtrack(n, perm, used)
}
