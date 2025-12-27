🧩 Tetris Backtracking Solver (Go)

This project solves a Tetris-like puzzle using a backtracking algorithm written in Go.
It reads shapes from a text file and tries to place them on the smallest possible square board without overlaps.

Each shape is represented using # and converted internally to unique letters (A, B, C, ...).

🚀 Features

Reads shapes from a text file

Automatically assigns a letter to each shape

Uses backtracking to find a valid placement

Dynamically increases board size until a solution is found

Prints the final board to the terminal

Safe execution (no crashes)

📂 Input Format

Shapes are separated by empty lines

# represents a block

. represents empty space

Example:

##..
.##.

####
....

🧠 How It Works

Read the input file

Split shapes by blank lines

Store each shape with a unique letter

Start with a 4 x 4 board

Try to place shapes one by one

If placement fails, backtrack

Increase board size until solved

🧱 Board Representation

The board is a slice of strings

Empty cells are .

Filled cells contain shape letters (A, B, C, ...)

Example output:

AA..
.ABB
.CBB
.CC.

🔄 Backtracking Logic

can_place → checks if a shape fits

place_shape → places the shape on the board

remove_shape → removes the shape (backtracking)

solve_tetris → recursive solver

Stops as soon as a valid solution is found ✅

▶️ Usage
go run main.go sample.txt


If the input is valid, the solved board will be printed.

⚠️ Error Handling

Missing arguments

File read errors

Invalid placements

The program exits safely with an error message when needed.

🧠 Concepts Covered

Backtracking

Recursion

File parsing

2D grid logic

Go slices and strings

Algorithmic problem solving
