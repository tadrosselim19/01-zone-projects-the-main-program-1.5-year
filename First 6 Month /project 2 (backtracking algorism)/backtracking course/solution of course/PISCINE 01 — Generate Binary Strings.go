package main

import "fmt"

func Generate_Binary_Strings(bit_count int,s string) {
	if len(s) == bit_count{
		fmt.Println(s)
		return
	}
	Generate_Binary_Strings(bit_count,s+"0")
	Generate_Binary_Strings(bit_count,s+"1")
}
func main() {
	bit_count := 3
	Generate_Binary_Strings(bit_count,"")
}