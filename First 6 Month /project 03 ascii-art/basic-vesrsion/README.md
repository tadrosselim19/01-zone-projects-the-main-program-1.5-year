# 🧱 ASCII Art Generator (Go)

This project generates ASCII-art text from a given font file.  
Each printable ASCII character (from 32 to 126) is stored in the file, separated by double newlines.  
When you run the program, it reads that font file and prints your word in styled ASCII art.

---

## 🚀 Features

- Reads custom ASCII font files (like `standard.txt`, `shadow.txt`, etc.)
- Supports multiline words using `\n`
- Prints any combination of printable ASCII characters
- Handles spacing and alignment properly
- Includes clean code structure and helper functions for splitting data

---

## 🧩 Project Structure

```
ascii-art/
│
├── main.go                    # Main program
├── ascii_test.go              # Unit tests
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
go run . standard.txt "ABC"
```

### Output
```
AABBCC
AABBCC
```

### Multiline Example
```bash
go run . standard.txt "HELLO\nTHERE"
```

### Output
```
<HELLO in ASCII Art>
<THERE in ASCII Art>
```

---

## 🧰 Testing

You can run the included unit tests using:
```bash
go test -v
```

This will test:
- Splitting font data correctly
- Handling newline sequences in text input
- ASCII rendering behavior

---

## 🧱 Functions Overview

| Function | Description |
|-----------|-------------|
| `split_and_remove_extra_newline()` | Splits ASCII font data into a map of runes and their ASCII patterns |
| `splid_by_newline_know_location()` | Splits the input text by `\n` markers while preserving blank lines |
| `main()` | Reads file and renders the input word using the ASCII map |

---

## 🧑‍💻 Author

**Tadros Selim**  
Created as part of learning and practicing Go language projects.

---

## 📜 License

This project is open-source and free to use for educational and non-commercial purposes.
