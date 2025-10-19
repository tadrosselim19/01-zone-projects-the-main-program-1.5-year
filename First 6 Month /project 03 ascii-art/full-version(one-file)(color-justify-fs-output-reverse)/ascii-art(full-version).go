package main

import (
	"fmt"
	"os"
	"strings"
)

// function that map what inside file
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

	return list_of_ascii_in_file
}

// function to mange edge case in multi \n added by user
func splid_by_newline_know_location(str string) [][]string {
	final := [][]string{}
	prefinal := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i:i+2] == "\\n" {
			prefinal = append(prefinal, word)
			final = append(final, prefinal)
			prefinal = []string{}
			i++
			word = ""
		} else if str[i] == ' ' {
			prefinal = append(prefinal, word, " ")
			word = ""

		} else {
			word += string(str[i])
		}

	}
	if word != "" {
		prefinal = append(prefinal, word)
		final = append(final, prefinal)
	}
	return final
}

// function that mangemnet if their order from user or not
func find_and_define_order(argument []string) (string, string, string, string, string) {
	if len(argument) > 4 {
		println("too many argument")
		return "", "", "", "", ""
	} else if len(argument) < 0 {
		println("missing comment line argument")
		return "", "", "", "", ""
	}

	if argument[0][:8] == "--color=" {
		if len(argument) == 4 {
			return "color", argument[0][8:], argument[1], argument[2], argument[3]
		} else if len(argument) == 3 {
			return "color", argument[0][8:], "", argument[1], argument[2]

		}

	} else if argument[0][:9] == "--output=" {
		out_file, err := os.Create(argument[0][9:])
		if err != nil {
			fmt.Println("system cant create this file")
			return "", "", "", "", ""
		}
		output_file_name := argument[0][9:]
		defer out_file.Close()
		return "output", output_file_name, "", argument[1], argument[2]

	} else if argument[0][:10] == "--reverse=" {
		return "reverse", argument[0][10:], "", "", argument[1]
	} else if argument[0][:8] == "--align=" {
		return "align", argument[0][8:], "", argument[1], argument[2]
	} else if len(argument) == 2 {
		return "", "", "", argument[0], argument[1]
	}
	return "", "", "", "", ""

}

// function of color
func color_order_check(color string) string {
	var colors = map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"reset":   "\033[0m",
	}
	if val, ok := colors[color]; ok {
		return val
	}
	fmt.Printf("This %s color isnt supported here", color)
	return ""
}

func reverse_assci_sparator_lines(str string) []string {
	length := len(str)
	final := []string{}
	line := ""

	for i := 0; i < length; i++ {
		if str[i] == '\n' {
			final = append(final, line)
			line = ""
			continue
		}
		line += string(str[i])
	}
	if line != "" {
		final = append(final, line)
	}
	return final
}
func reverse_find_location_of_sperator(lines []string) []int {
	length := len(lines[0])
	final := []int{}

	for i := 0; i < length; i++ {
		line_increment := 0
		found_sperator := false
		if lines[line_increment][i] == ' ' {
			for line_increment < len(lines)-1 {
				line_increment++
				if lines[line_increment][i] != ' ' {
					found_sperator = false
					break
				} else {
					found_sperator = true
				}
			}

		}
		if found_sperator == true {
			final = append(final, i)
		}
	}

	return final
}
func reverse_assci(file string, list map[rune]string) {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("file not exist on system")
		return
	}
	lines := reverse_assci_sparator_lines(string(input))
	stoper_indexs := reverse_find_location_of_sperator(lines)

	final := ""
	stared := 0
	for j := 0; j < len(stoper_indexs); j++ {
		charter := ""
		count_lines := 0

		for i := stared; count_lines < 8; {
			if i == stoper_indexs[j] {
				count_lines++
				if count_lines != 7 {
					charter += " \n"
				} else {
					charter += "\n"
				}
				i = stared
			} else {
				charter += string(lines[count_lines][i])
				i++
			}
		}

		for key, value := range list {
			cleanValue := strings.TrimSpace(value)
			cleanCharter := strings.TrimSpace(charter)
			if cleanValue == cleanCharter {
				final += string(key)

			}

		}
		stared = stoper_indexs[j] + 1

	}

	if len(final) == len(stoper_indexs) {
		fmt.Println(final)
		return
	} else {
		fmt.Println("the file contain charchter not in file you given")
	}

}
func main() {
	if len(os.Args) < 2 {
		return
	}
	// extract all needed optsion
	order, action, detials, text, file_name := find_and_define_order(os.Args[1:])
	if text == "" && file_name == "" && order != "reverse" {
		fmt.Println("Usage: go run . (order)option text file.txt")
		return
	}

	// opening file of assci and extrat input
	file, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("can not open file")
		return
	}
	input := string(file)
	list_of_ascii := split_and_remove_extra_newline(input)
	//word := os.Args[2]
	word_array_large := splid_by_newline_know_location(text)

	color := ""
	color_word := ""
	output_file_name := ""
	direction_of_align := ""
	var output_file *os.File
	var err1 error
	if order == "color" {
		color += action
		color_word += detials
	} else if order == "output" {
		output_file_name += action
	} else if order == "align" {
		direction_of_align += action
	} else if order == "reverse" {
		reverse_assci(action, list_of_ascii)
		return
	}

	//1 - mange color-assci
	if color != "" {
		color_code := color_order_check(color)
		if color_code == "" {
			return
		}
		color = color_code
	}

	// 2- mange output
	if output_file_name != "" {
		output_file1, err := os.Create(output_file_name)
		if err != nil {
			fmt.Println("cant create file")
			return
		}
		defer output_file1.Close()
		output_file, err1 = os.OpenFile(output_file_name, os.O_APPEND, 0644)
		if err1 != nil {
			panic(err)
		}

	}

	// 3- mange align
	width := 8 * len(text)
	width_control := 0
	if action == "right" || action == "left" || action == "center" {
		if action == "right" {
			width_control += 1
		} else if action == "left" {
			width_control += 3
		} else if action == "center" {
			width_control += 2
		}

	}

	// printing
	for _, word_array := range word_array_large {

		// prepare started_index same as you already do earlier
		last_index_last_word := 0
		started_index := [][]int{}
		for i := 0; i < len(word_array); i++ {
			size := make([]int, len(word_array[i]))
			started_index = append(started_index, size)
			if i == len(word_array)-1 {
				last_index_last_word += len(word_array[i])
			}
		}

		// main loop: build one printed row (across all words and letters) per iteration
		for {
			line_g := ""
			allDone := true // assume done until we find a letter that still has data
			switch action {
			case "right":
				line_g = "|" + strings.Repeat(" ", width*2) + line_g
			case "center":
				line_g = "|" + strings.Repeat(" ", width) + line_g
			case "left":
				line_g = "|" + line_g
			case "justify":
				line_g = "|" + line_g
			}

			for i := 0; i < len(word_array); i++ {
				line_w := ""

				// start color prefix for the whole word if needed
				if word_array[i] == color_word || (color_word == "" && color != "") {
					line_w += color
				} else {
					// don't add a reset here; we will append reset after the full printed line
					line_w += ""
				}

				for j := 0; j < len(word_array[i]); j++ {
					letter_str := list_of_ascii[rune(word_array[i][j])]
					start := started_index[i][j]
					line := ""

					// if we've already consumed this letter fully, append a gap
					if start >= len(letter_str) {
						// preserve width by adding blank column(s)
						line = "      "
						line_w += line
						continue
					}

					// read characters until next newline or end
					foundNewline := false
					for k := start; k < len(letter_str); k++ {
						if letter_str[k] == '\n' {
							// move start to byte after '\n' for next row
							started_index[i][j] = k + 1
							foundNewline = true
							break
						}
						line += string(letter_str[k])
					}

					// if no newline found, we've consumed to the end — advance start to end
					if !foundNewline {
						started_index[i][j] = len(letter_str)
					}

					// if this letter still has more to print in future iterations, we're not done
					if started_index[i][j] < len(letter_str) {
						allDone = false
					}

					if line == "" {
						line = "      "
					}
					line_w += line
				}

				// append built word part to global line
				line_g += line_w
			}

			switch action {
			case "right":
				line_g = line_g + "|"
			case "center":
				line_g =  line_g + strings.Repeat(" ", width) + "|"
			case "left":
				line_g = line_g + strings.Repeat(" ", width*2) + "|"
			case "justify":
				line_g = line_g + "|"
			}
			// print/write the assembled line (always print the last line too)
			if output_file_name != "" {
				_, _ = output_file.WriteString(line_g + "\n")
			} else {
				// reset color for safety after the printed line
				fmt.Println(line_g + "\033[0m")
			}

			// if all letters across all words are fully consumed, stop
			if allDone {
				break
			}
		}
		println()
		// -------------------------
	}
	if width_control == 3 {
		fmt.Print(strings.Repeat(" ", width*2))
	}
	if width_control == 2 {
		fmt.Print(strings.Repeat(" ", width))
	}

	defer output_file.Close()
	return

}

// 32 -> 126
