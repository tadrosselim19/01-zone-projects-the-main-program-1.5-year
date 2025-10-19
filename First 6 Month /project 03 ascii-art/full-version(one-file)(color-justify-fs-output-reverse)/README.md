# 🧱 ASCII Art Generator (Go)

This project generates ASCII-art text from a given font file.  
Each printable ASCII character (from 32 to 126) is stored in the file, separated by double newlines.  
When you run the program, it reads that font file and prints your word in styled ASCII art.

---

## 🚀 Features

- Reads custom ASCII font files (like `standard.txt`, `shadow.txt`, etc.)
- Supports multiline words using `\n`
- Prints any combination of printable ASCII characters
- Handles spacing and alignment properly (`left`, `right`, `center`, `justify`)
- Supports colored output
- Can save ASCII art to a file
- Can reverse ASCII art back into text
- Includes clean code structure and helper functions for splitting data

---

## 🧩 Project Structure

```
ascii-art/
│
├── main.go                    # Main program
├── standard.txt               # Example ASCII font file
└── README.md                  # This file
```

---

## ⚙️ How It Works

1. **Font file parsing:**  
   The `split_and_remove_extra_newline()` function splits the font file into blocks, each representing a printable ASCII character.

2. **Word rendering:**  
   The program loops through each character of the word, printing its ASCII art line by line.

3. **Multiline words:**  
   You can use `\n` in your input to break lines (e.g. `HELLO\nTHERE`).  
   The `splid_by_newline_know_location()` function handles splitting and spacing for each line.

4. **Color and alignment:**  
   You can apply ANSI colors and align text using `--color=` and `--align=` options.

5. **Reverse ASCII art:**  
   The program can also read an ASCII art file and reconstruct the original text with `--reverse=`.

---

## 🧪 Example Usage

### Example Input File (`standard.txt`)

```
AA
AA

BB
BB

CC
CC
```

### Command

```bash
go run main.go --align=left --color=green "ABC" standard.txt
```

### Output

```
AABBCC
AABBCC
```

### Multiline Example

```bash
go run main.go "HELLO\nTHERE" standard.txt
```

### Output

```
<HELLO in ASCII Art>
<THERE in ASCII Art>
```

### Reverse ASCII

```bash
go run main.go --reverse=art.txt standard.txt
```

---

## 🧰 Command-line Options

| Option          | Description                                           | Example                          |
|-----------------|-------------------------------------------------------|----------------------------------|
| `--color=COLOR` | Set the color of the text (`red`, `green`, etc.)     | `--color=red`                    |
| `--output=FILE` | Save ASCII art to a file                              | `--output=art.txt`               |
| `--align=TYPE`  | Text alignment: `left`, `right`, `center`, `justify`| `--align=center`                 |
| `--reverse=FILE`| Convert ASCII art back to text                        | `--reverse=art.txt`              |

---

## 🧱 Functions Overview

| Function | Description |
|-----------|-------------|
| `split_and_remove_extra_newline()` | Splits ASCII font data into a map of runes and their ASCII patterns |
| `splid_by_newline_know_location()` | Splits the input text by `\n` markers while preserving blank lines |
| `find_and_define_order()` | Parses command-line arguments to determine actions (color, output, align, reverse) |
| `color_order_check()` | Returns ANSI color code for supported colors |
| `reverse_assci_sparator_lines()` | Splits ASCII art into individual lines for reverse processing |
| `reverse_find_location_of_sperator()` | Detects vertical separators between characters in ASCII art |
| `reverse_assci()` | Converts ASCII art back into normal text |
| `main()` | Reads the font file, parses arguments, and renders ASCII art with alignment, color, and output options |

---

## 🧑‍💻 Author

**Tadros Selim**  
Created as part of learning and practicing Go language projects.

---

## 📜 License

This project is open-source and free to use for educational and non-commercial purposes.

