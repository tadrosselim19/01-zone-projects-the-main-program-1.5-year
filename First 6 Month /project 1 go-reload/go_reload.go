package main

import (
	"fmt"
	"os"
	"strconv"
	
)

func splid_input(str string) []string {
	final := []string{}
	sentense := ""
	len := len(str)
	bracket := ""

	for i := 0; i < len; i++ {
		if str[i] == '(' {
			bracket += string(str[i])
		}
		if str[i] == ')' {
			bracket = ""
		}
		if str[i] != ',' {
			sentense += string(str[i])
		} else if str[i] == ',' && bracket == "" {
			final = append(final, sentense)
			sentense = ""
		}
		// }else if str[i] == ',' && bracket != "" {
		// 	sentense += string(str[i])
		// }
	}
	if sentense != "" {
		final = append(final, sentense)
	}

	return final
}

func to_decimal(num string, action string) string {
	
	dicimal := 0
	if action == "bin" {
		power := 1
		for i := len(num) - 1; i >= 0; i-- {
			if num[i] == '0' || num[i] == '1' {
				dicimal += int((num[i] - '0')) * power
				power *= 2
			} else {
				return ""
			}
		}
	} else if action == "hex" {
		power := 1
		for i := len(num) - 1; i >= 0; i-- {
			if num[i] >= '0' && num[i] <= '9' {
				dicimal += int((num[i] - '0')) * power
				power *= 16
			} else if num[i] >= 'a' && num[i] <= 'f' {
				dicimal += (int(num[i] - 'a') + 10) * power
				power *= 16
			} else if num[i] >= 'A' && num[i] <= 'F' {
				dicimal += (int(num[i] - 'A') + 10) * power
				power *= 16
			} 
		}
	}

	return strconv.Itoa(dicimal)
}


func change_case(sentense string, action string, repate int) string {
	str := []byte(sentense)
	len := len(str)
	if action == "cap" {
		for i := len - 2; i >= 0 && repate != 0; i-- {
			if i-1 < 0 && str[i] >= 'a' && str[i] <= 'z' {
				str[i] = str[i] - 32
				break
			}
			if i-1 >= 0 && str[i-1] == ' ' && str[i] >= 'a' && str[i] <= 'z' {
				str[i] = str[i] - 32
				repate--
			}
		}
	}
	if action == "up" {
		for i := len - 1; i >= 0 && repate != 0; i-- {
			if i-1 >= 0 && str[i-1] == ' ' {
				for j := i; j < len && str[j] != ' '; j++ {
					if str[j] >= 'a' && str[j] <= 'z' {
						str[j] = str[j] - 32
					}
				}
				repate--
			}
		}

	}

	if action == "low" {
		for i := len - 1; i >= 0 && repate != 0; i-- {
			if i-1 >= 0 && str[i-1] == ' ' {
				for j := i; j < len && str[j] != ' '; j++ {
					if str[j] >= 'A' && str[j] <= 'Z' {
						str[j] = str[j] + 32
					}
				}
				repate--
			}
		}

	}

	return string(str)
}

func a_an_case(str string) string {
	final := ""
	for i := 0; i < len(str); i++ {
		if i+2 < len(str) && str[i:i+2] == "a " && 
			(str[i+2] == 'a' || str[i+2] == 'A' || str[i+2] == 'E' || str[i+2] == 'e' || 
			 str[i+2] == 'I' || str[i+2] == 'i' || str[i+2] == 'O' || str[i+2] == 'o' || 
			 str[i+2] == 'U' || str[i+2] == 'u' || str[i+2] == 'h' || str[i+2] == 'H') {
			final += "an "
			i++
			continue
		}
		if i+2 < len(str) && str[i:i+2] == "A " && 
			(str[i+2] == 'a' || str[i+2] == 'A' || str[i+2] == 'E' || str[i+2] == 'e' || 
			 str[i+2] == 'I' || str[i+2] == 'i' || str[i+2] == 'O' || str[i+2] == 'o' || 
			 str[i+2] == 'U' || str[i+2] == 'u' || str[i+2] == 'h' || str[i+2] == 'H') {
			final += "An "
			i++
			continue
		}
		final += string(str[i])
	}
	return final
}

func punctuations_modification(str string) string {
	final := ""
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == ' ' && (str[i+1] == ',' || str[i+1] == ':' || str[i+1] == '.' || str[i+1] == '!' || str[i+1] == '?' || str[i+1] == ';') {
			continue
		}

		final += string(str[i])
		if i+1 < len(str) && (str[i] == ',' || str[i] == ':' || str[i] == '.' || str[i] == '!' || str[i] == '?' || str[i] == ';') && (str[i+1] == ',' || str[i+1] == ':' || str[i+1] == '.' || str[i+1] == '!' || str[i+1] == '?' || str[i+1] == ';') {
			continue
		}
		if i+1 < len(str) && (str[i] == ',' || str[i] == ':' || str[i] == '.' || str[i] == '!' || str[i] == '?' || str[i] == ';') && str[i+1] != ' ' {
			final += " "
		}
	}

	return final
}

func qutation_modification(str string) string {
	final := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '\'' && i+1 < len(str) && str[i+1] == ' ' {
			final += "'"
			i++ // skip the space
			continue
		}
		if i+1 < len(str) && str[i] == ' ' && str[i+1] == '\''   {
			final += "'"
			i++
			continue
		}
		
		final += string(str[i])

	}
	return final
}

func define_order(order string) (string, int) {
	final_order := ""
	final_repate := ""
	for i := 0; i < len(order); i++ {
		if order[i] >= 'a' && order[i] <= 'z' {
			final_order += string(order[i])
		}
		if order[i] >= '0' && order[i] <= '9' {
			final_repate += string(order[i])
		}
	}
	repete, err := strconv.Atoi(final_repate)
	if err != nil {
		return final_order, 1
	}
	return final_order, repete

}

func space_mangment(str string) string {
	final := ""
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 && str[i] == ' ' {
			break
		}
		if i+1 < len(str) && str[i] == ' ' && str[i+1] == ' ' {
			continue
		}
		final += string(str[i])
	}
	return final
}

func add_result_to_file(str string, file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("cant create file")
		return
	}
	file.WriteString(str)
	file.Close()
	return
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <source_file> <result_file>")
		return
	}

	source_file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("cant allocate the file")
		return
	}

	//mod_file = punctuations_modification(mod_file)
	mod_file := string(source_file)
	len_file := len(mod_file)

	final := ""
	sentense := ""
	bracket := false
	order := ""
	repeate := 1
	for i := 0; i < len_file; i++ {
		if mod_file[i] == '(' {
			bracket = true
		}
		if bracket == false {
			sentense += string(mod_file[i])
		} else {
			order += string(mod_file[i])

			if mod_file[i] == ')' {
				bracket = false

				order, repeate = define_order(order)
				if order == "cap" || order == "up" || order == "low" {
					sentense = change_case(sentense, order, repeate)
					final += sentense
					sentense = ""
					order = ""

				} else if order == "hex" || order == "bin" {
					num_str := ""
					num_final := ""
					flag := true
					for j := len(sentense) - 1; j >= 0; j-- {
						if sentense[j] != ' ' {
							flag = false
							num_str += string(sentense[j])
						}
						if sentense[j] == ' ' && flag == false {
							break
						}
					}

					for j := len(num_str) - 1; j >= 0; j-- {
						num_final += string(num_str[j])
					}

					sentense = sentense[:len(sentense)-(len(num_final)+1)]
					num_to_be_add := to_decimal(num_final, order)
					sentense += num_to_be_add
					final += sentense

					sentense = ""
					order = ""
				} else {
					fmt.Println("you want to run a wrong order")
					return
				}
			}

		}

	}
	if sentense != "" {
		final += sentense
	}
	final = space_mangment(final)
	final = punctuations_modification(final)
	final = qutation_modification(final)
	final = a_an_case(final)

	add_result_to_file(final, os.Args[2])

}
