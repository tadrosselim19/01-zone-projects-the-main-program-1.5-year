🧩 Tetris Optimizer – Version 1 (Basic Backtracking)
📌 Overview

This project is a basic Tetris optimizer written in Go, inspired by the 01-school tetris-optimizer problem.

The goal is to place a set of tetromino shapes into the smallest possible square board using a backtracking algorithm.

This first version focuses on:

Clean input parsing

Simple backtracking placement

Incrementally increasing board size

No rotations (yet)

🛠️ Features

Reads tetromino shapes from a text file

Assigns each shape a unique letter (A, B, C, ...)

Uses backtracking to try all placements

Starts from a 4×4 board and increases size until a solution is found

Prints the first valid solution

📂 Input Format

Each shape is 4 lines

# represents a block

. represents empty space

Shapes are separated by an empty line

Example input file:
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

▶️ How It Works

The file is read and split into shapes

Each # is replaced with a unique letter

A square board is created and initialized with .

Shapes are placed one by one using backtracking

If placement fails, the algorithm backtracks

The board size increases until a solution is found

🚀 Usage
go run main.go sample.txt

🧠 Core Algorithm

Backtracking

Try every (x, y) position for each shape

Check validity using bounds and collision detection

Place → recurse → remove if failed

Key functions:

can_place

place_shape

remove_shape

solve_tetris

📌 Limitations (Version 1)

❌ No shape rotations

❌ No shape normalization (top-left alignment)

❌ Only finds the first valid solution

❌ Not optimized for speed

These limitations are intentional and addressed in later versions.

🧪 Example Output
AABB
A..B
CCCB
....

📈 Next Versions

Planned improvements in later versions:

Shape rotation handling

Shape trimming / normalization

Reduced branching and optimizations

Faster solving for larger inputs

👨‍💻 Author

Tadros Selim
Go | Algorithms | Backtracking | 01 Talent Program
