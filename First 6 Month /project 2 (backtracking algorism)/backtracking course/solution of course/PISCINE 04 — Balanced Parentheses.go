package main

import "fmt"

func Balanced_Parentheses(n int,s string, open string,close string) {
	if len(s) == n*2{
		fmt.Println(s)
		return
	}
	
	if len(open) < n{
		open += "("
		s+="("
		Balanced_Parentheses(n,s,open,close)
		open=open[:len(open)-1]
		s=s[:len(s)-1]
	}
	
	if len(close) < len(open){
		close+=")"
		s+=")"
		Balanced_Parentheses(n,s,open,close)
		close=close[:len(close)-1]
		s=s[:len(s)-1]
		
	}
	// fmt.Println(open)
	// fmt.Println(close)
	// fmt.Println(s)	
	//s = s[:len(s)-n]
	
}
func Balanced(n int, s string, open int, close int) {
	if len(s) == n*2 {
		fmt.Println(s)
		return
	}

	// Can add "(" if we have not used all
	if open < n {
		Balanced(n, s+"(", open+1, close)
	}

	// Can add ")" only if we have more "(" than ")"
	if close < open {
		Balanced(n, s+")", open, close+1)
	}
}

func main() {
	Parentheses_num := 3
	Balanced_Parentheses(Parentheses_num, "","","")
	println("\n-------\n")
	Balanced(Parentheses_num,"",0,0)
}
