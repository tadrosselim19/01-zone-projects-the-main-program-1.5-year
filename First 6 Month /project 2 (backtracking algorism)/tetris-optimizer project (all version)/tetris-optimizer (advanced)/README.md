🧩 Tetris Optimizer (Go) – Rotation Version

This project is a Tetris (Fillit-style) optimizer written in Go.
It reads tetromino shapes from a text file and places them into the smallest possible square board.

This version supports shape rotation and uses an improved anchor-based placement algorithm.

🚀 Features

Reads tetromino shapes from a text file

Assigns each shape a unique letter (A, B, C, ...)

Supports Windows (\r\n) and Unix (\n) line endings

Uses backtracking to find a valid placement

Supports 90° rotations for each shape

Uses the first block as a reference point for accurate placement

Automatically increases board size until a solution is found

🧩 Project Structure
tetris-optimizer/
│
├── main.go        # Main solver and backtracking logic
├── sample.txt     # Input shapes file
└── README.md      # This file

⚙️ How It Works

Shape parsing:
The split_to_define_shape() function reads the file and converts each # shape into a letter-based shape (A, B, C, …).

Anchor-based placement:
The first non-. block in each shape is used as a reference point.
This allows shapes to be placed correctly even if they are not top-left aligned.

Rotation handling:
Each shape is rotated up to 4 times (0°, 90°, 180°, 270°) using rotate_90().

Backtracking solver:
The solver places shapes one by one.
If placement fails, it backtracks and tries another position or rotation.

🧪 Example Input File
...#
...#
...#
...#

....
....
....
####


Each shape:

Uses # for blocks

Uses . for empty space

Is separated by an empty line

▶️ Example Usage
go run main.go sample.txt

📤 Example Output
AABB
A..B
CCCB
....


Each letter represents one tetromino.

🧠 Algorithm

Backtracking

Recursive depth-first search

Collision and boundary checks

Rotation-based branching

Key functions:

can_place

place_shape

remove_shape

rotate_90

solve_tetris

🧱 Functions Overview
Function	Description
split_to_define_shape()	Parses input file and creates shape list
rotate_90()	Rotates a shape 90 degrees clockwise
can_place()	Checks if a shape can be placed at a position
place_shape()	Places a shape on the board
remove_shape()	Removes a shape (backtracking)
solve_tetris()	Recursive backtracking solver
make_bord()	Creates an empty board
❌ Limitations

No shape normalization (trimming empty rows/columns)

No symmetry optimization (duplicate rotations)

Stops at the first valid solution

Performance not optimized for large inputs

These are improved in future versions.

🧑‍💻 Author

Tadros Selim
Go • Algorithms • Backtracking • Tetris Optimizer
01 Talent Program

📜 License

This project is open-source and free to use for learning and educational purposes.
