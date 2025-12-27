package main

import "fmt"

func subsets_desc(arr []int, index int, current []int) {
    // Base case: reached the end
    if index == len(arr) {
        fmt.Println(current)
        return
    }

    // include arr[index]
    current = append(current, arr[index])
    subsets_desc(arr, index+1, current)

    // backtrack and exclude arr[index]
    current = current[:len(current)-1]
    subsets_desc(arr, index+1, current)
}

// Backwards version: consider elements from right -> left
func subsets_acda(arr []int, index int, current []int) {
    // Base case: no more elements to consider
    if index < 0 {
        fmt.Println(current)
        return
    }

    // 1) skip arr[index]
    subsets_acd(arr, index-1, current)

    // 2) include arr[index]
    current = append(current, arr[index])
    subsets_acd(arr, index-1, current)

    // backtrack (not strictly necessary here because `current` is a local slice value,
    // but good style to show the pattern)
    // current = current[:len(current)-1]
}
func subsets_acd(arr []int, index int, current []int) {
    if index == 0 {
        fmt.Println(current)
        return
    }

    // 1) exclude arr[index-1]
    subsets_acd(arr, index-1, current)

    // 2) include arr[index-1] BUT prepend so order remains increasing
    // newCurrent = [ arr[index-1] ] + current
    newCurrent := make([]int, 1+len(current))
    newCurrent[0] = arr[index-1]
    copy(newCurrent[1:], current)

    subsets_acd(arr, index-1, newCurrent)
}

func main() {
    arr := []int{1, 2, 3}
    subsets_desc(arr, 0, []int{})
    fmt.Println("----")
    subsets_acd(arr, len(arr), []int{})
}
