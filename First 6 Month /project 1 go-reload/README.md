🧠 Text Formatter CLI (Go)

A command-line tool written in Go that reads a text file, applies a series of embedded text transformation commands, and writes a clean, formatted version to an output file.

This project was developed as part of the 01 Talent Program, focusing on string manipulation, text parsing, and file handling in Go.

🚀 Features

Text Transformations

(cap): Capitalize specific words from the end

(up): Convert words to uppercase

(low): Convert words to lowercase

Number Conversion

(hex): Converts the last hexadecimal number to decimal

(bin): Converts the last binary number to decimal

Grammar & Punctuation

Fixes spacing around punctuation marks

Adjusts single quotes placement

Automatically replaces “a” → “an” before vowels

File Handling

Reads input from a text file

Writes output to a result file

Manages spaces and formatting automatically

🧩 Example

Input (input.txt):
Harold Wilson:' I am an optimist, but an optimist who carries a raincoat.'
A 1a (hex) cat has 10 (bin) lives.

Command:
go run main.go input.txt output.txt

Output (output.txt):
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
A 26 cat has 2 lives.

⚙️ How to Use

Create an input text file (for example, input.txt).

Run the program:
go run main.go input.txt output.txt

Check your formatted output in output.txt

🧰 Functions Overview
Function	Description
to_decimal()	Converts hexadecimal or binary to decimal

change_case()	Changes capitalization, upper or lower case

a_an_case()	Fixes usage of “a” / “an” before vowels

punctuations_modification()	Adds or removes spaces near punctuation marks

qutation_modification()	Corrects spacing around single quotes

space_mangment()	Removes unnecessary spaces

add_result_to_file()	Saves the final output text to a file

🛠 Requirements

Go 1.20 or newer

Works on Linux, macOS, or Windows

📦 Project Structure

main.go
input.txt
output.txt
README.md

👤 Author

Tadros Selim
Participant – 01 Talent Program
Focus: Go language, logic, and clean code
