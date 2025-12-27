🧩 Tetris Optimizer (Go)

This project solves a Tetris packing problem using a backtracking algorithm.
It reads tetromino shapes from a file and tries to fit all of them into the smallest possible square board.
Each shape is labeled with a unique capital letter (A, B, C, …) and printed once a valid solution is found.

🚀 Features

Reads tetromino shapes from a text file

Supports both \n and \r\n line endings

Automatically assigns letters to shapes

Uses backtracking to explore valid placements

Starts from a 4×4 board and grows if needed

Prints the first valid solution found

🧩 Project Structure
tetris-optimizer/
│
├── main.go          # Main program and solver
├── sample.txt       # Example input file
└── README.md        # This file

⚙️ How It Works

Input parsing:
The split_to_define_shape() function reads the file, splits shapes by blank lines, and converts # into letters.

Board creation:
The make_bord() function creates an empty square board filled with dots (.).

Backtracking solver:
The solve_tetris() function tries to place each shape on the board.
If placement fails, it removes the shape and tries another position.

Board resizing:
The solver starts with size 4 and increases the board size until a solution is found.

🧪 Example Usage
Example Input File (sample.txt)
...#
...#
...#
...#

....
....
....
####

Command
go run main.go sample.txt

Output
AABB
A..B
CCCB
....

🧠 Algorithm

Backtracking

Depth-first search

Collision and boundary checking

Recursive placement and removal

Core functions:

can_place()

place_shape()

remove_shape()

solve_tetris()

🧱 Functions Overview
Function	Description
split_to_define_shape()	Parses the input file and builds shape data
can_place()	Checks if a shape fits at a given position
place_shape()	Places a shape on the board
remove_shape()	Removes a shape during backtracking
solve_tetris()	Recursive backtracking solver
make_bord()	Creates an empty square board

⚠️ Limitations

No shape rotations

No shape trimming

No performance optimizations

Stops at the first valid solution

These are handled in later versions of the project.

🧑‍💻 Author

Tadros Selim
Built while practicing Go, algorithms, and backtracking techniques.

📜 License

This project is open-source and free to use for educational and non-commercial purposes.
