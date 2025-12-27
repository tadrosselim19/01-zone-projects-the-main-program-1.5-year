package main

import "fmt"

func Password_Explorer(n int, s string) {
	if len(s) == n {
		if s != "xx"{
			fmt.Println(s)

		}
		return
	}
	for i := 'a' ; i <= 'z';i++{
		s += string(i)
		Password_Explorer(n,s)
		s = s[:len(s)-1]
	}
	
}

func main() {
	ch_num := 2
	Password_Explorer(ch_num, "")
}
