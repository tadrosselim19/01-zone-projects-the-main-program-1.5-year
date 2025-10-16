package main

import (
	"fmt"
	"os"
)

//import "fmt"

func split_and_remove_extra_newline(str string) map[rune]string {
	list_of_ascii_in_file := map[rune]string{}
	length := len(str)
	ascii_ch := ""
	ascii_num := ' '

	for i := 0; i < length; i++ {
		if str[i] != '\n' {
			ascii_ch += string(str[i])
			continue
		}
		if i+1 < length && str[i] == '\n' && str[i+1] != '\n' {
			ascii_ch += string(str[i])
		} else if i+1 < length && str[i] == '\n' && str[i+1] == '\n' {
			if ascii_ch != "" {
				i++
				list_of_ascii_in_file[ascii_num] = ascii_ch
				ascii_num++
				ascii_ch = ""
			}

		}
	}
	// check of function working
	// for _ ,i := range("hello"){
	// 	fmt.Println(list_of_ascii_in_file[i])
	// }

	return list_of_ascii_in_file
}
func splid_by_newline_know_location(str string) []string {
	final := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i:i+2] == "\\n" {
			final = append(final, word, "")
			i++
			word = ""
		} else {
			word += string(str[i])
		}

	}
	if word != "" {
		final = append(final, word)
	}
	return final
}
func main() {
	if len(os.Args) != 3 {
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("can not open file")
		return
	}
	input := string(file)
	list_of_ascii := split_and_remove_extra_newline(input)
	//word := os.Args[2]
	word_array := splid_by_newline_know_location(os.Args[2])

	for _, word := range word_array {
		if word ==""{
			println()
			continue
		}
		// array for determin when i stop and begin
		stared_index := make([]int, len(word))

		done := false

		for !done {
			done = true
			for i := 0; i < len(word); i++ {
				letter_str := list_of_ascii[rune(word[i])]
				start := stared_index[i]
				line := ""

				for j := start; j < len(letter_str); j++ {
					if letter_str[j] == '\n' {
						stared_index[i] = j + 1
						break
					}
					line += string(letter_str[j])
					if j == len(letter_str)-1 {
						stared_index[i] = len(letter_str)
					}
				}

				if stared_index[i] < len(letter_str) {
					done = false
				}

				if line == "" {
					fmt.Print("      ") // spacing for empty line
				} else {
					fmt.Print(line)
				}
			}
			if !done {
				fmt.Println()
			}
		}
	}

}

// 32 -> 126
