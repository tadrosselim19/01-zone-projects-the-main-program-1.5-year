package main
import ( 
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
func GenerateASCII(text string ,src_file string) string{
	
	file, err := os.ReadFile(src_file)
	if err != nil {
		return "can not open file"
	}
	input := string(file)
	list_of_ascii := split_and_remove_extra_newline(input)
	//word := os.Args[2]
	word_array := splid_by_newline_know_location(text)


	final_result := ""
	for _, word := range word_array {
		if word == ""{
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
					// spacing for empty line
					final_result += "      "
				} else {
					final_result += line
				}
			}
			if !done {
				final_result += "\n"
			}
		}
	}
	return final_result
}

func split_text_before_coloring(text string , sub string)[]string{
	final := []string{}
	word := ""
	for i := 0 ; i < len(text) ; i++{
		if i+len(sub)<=len(text) && text[i:i+len(sub)] == sub{
			final = append(final, word)
			final = append(final, sub)
			word = ""
			i += len(sub)-1
		}else{
			word += string(text[i])
		}
	}
	if word != ""{
		final = append(final, word)
	}	
	return final

}