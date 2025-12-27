# 🧩 Backtracking Piscine — Go Edition

A **project-based, Piscine-style course** to master **Backtracking algorithms using Go**.

This course is inspired by **42/01 Piscine problem sets** and is designed to take you from **basic recursion** all the way to a **Tetris Optimizer (10/10 difficulty)**.

---

## 🎯 Course Objectives

By the end of this course, you will be able to:

- Understand backtracking deeply (not just use it)
- Design recursive search trees
- Apply pruning and optimization techniques
- Solve board-based and constraint-satisfaction problems
- Build complex solvers (Sudoku, Shape Packing, Tetris Optimizer)
- Write clean, efficient Go backtracking code

---

## 🧠 What Is Backtracking?

Backtracking is an algorithmic technique that:

1. Tries a choice
2. Checks if it is valid
3. Continues exploring
4. Reverts the choice (undo)
5. Tries the next option

It is used for problems where **all possibilities must be explored**, but **invalid paths should be abandoned early**.

---

## 🏊 Piscine Rules

- Language: **Go (Golang)**
- No external libraries (standard library only)
- Recursion is mandatory where applicable
- Brute force without pruning will fail in later phases
- Output format must match exactly
- Order of output does **not** matter unless specified

---

# 📘 PHASE 1 — Backtracking Foundations (Difficulty: 2–4/10)

### Focus
- Recursion basics
- Choice exploration
- Undoing decisions

---

### 🟦 Piscine 01 — Binary Strings Generator (2/10)

**Task:**
Generate all binary strings of length `N`.

**Input:**
```
3
```

**Expected Output:**
```
000
001
010
011
100
101
110
111
```

---

### 🟦 Piscine 02 — Permutations of 1..N (3/10)

**Task:**
Print all permutations of numbers from `1` to `N`.

**Input:**
```
3
```

**Expected Output:**
```
1 2 3
1 3 2
2 1 3
2 3 1
3 1 2
3 2 1
```

---

### 🟦 Piscine 03 — Subset Generator (3/10)

**Task:**
Print all subsets of a given list.

**Input:**
```
1 2 3
```

**Expected Output (order may vary):**
```
[]
[1]
[2]
[3]
[1 2]
[1 3]
[2 3]
[1 2 3]
```

---

### 🟩 PHASE 1 PROJECT — Password Explorer (4/10)

**Task:**
Generate all strings of length `N` using `a–z`.

**Constraint (Trick):**
- Character `x` **cannot appear twice in a row**

---

# 📘 PHASE 2 — Constraints & Pruning (Difficulty: 4–6/10)

### Focus
- Validity checks
- Early pruning
- Constraint satisfaction

---

### 🟦 Piscine 04 — Balanced Parentheses (4/10)

**Input:**
```
3
```

**Expected Output:**
```
((()))
(()())
(())()
()(())
()()()
```

---

### 🟦 Piscine 05 — Sum ≤ K Constraint (4/10)

**Input:**
```
N = 4, K = 5
```

**Expected Output:**
```
[]
[1]
[2]
[1 2]
[3]
[1 3]
[4]
[1 4]
```

---

### 🟩 PHASE 2 PROJECT — Obstacle Grid Path Finder (6/10)

**Task:**
Find all paths from `(0,0)` to `(N-1,N-1)`.

**Trick:**
- Cells marked `T` are traps
- Stepping on `T` forces an immediate backward move

---

# 📘 PHASE 3 — Board Backtracking (Difficulty: 5–7.5/10)

### Focus
- 2D boards
- Mark / unmark logic
- Movement-based recursion

---

### 🟦 Piscine 06 — Maze Solver (5/10)

Find one valid path from `S` to `E`.

---

### 🟦 Piscine 07 — Knight’s Tour (7/10)

Find **one valid Knight’s Tour** on a 5×5 board.

---

### 🟩 PHASE 3 PROJECT — N-Queens + Symmetry Filter (7.5/10)

**Trick:**
- Reject mirror-symmetric solutions

---

# 📘 PHASE 4 — Optimization Backtracking (Difficulty: 6–8/10)

### Focus
- Best-solution tracking
- Branch & Bound
- Cost-based pruning

---

### 🟦 Piscine 08 — Minimum Path Sum (6/10)

Find the minimum-cost path using backtracking.

---

### 🟦 Piscine 09 — Minimal Palindrome Partition (7/10)

Split a string into the smallest number of palindromes.

---

### 🟩 PHASE 4 PROJECT — Forbidden Palindrome Partition (8/10)

**Trick:**
- Certain palindromes (`aa`, `bb`, `aba`) are forbidden

---

# 📘 PHASE 5 — Polyomino & Shape Packing (Difficulty: 6–8.5/10)

### Focus
- Shape placement
- Rotation generation
- Grid backtracking

---

### 🟦 Piscine 10 — Domino Tiling (6/10)

Determine if a board can be tiled with dominos.

---

### 🟦 Piscine 11 — Tetromino Rotations (6/10)

Generate all **unique rotations** of a tetromino.

---

### 🟩 PHASE 5 PROJECT — Mini Shape Packing (8.5/10)

**Trick:**
- One shape supports only **2 rotations**

---

# 📘 PHASE 6 — Advanced CSP (Difficulty: 7–9/10)

### Focus
- Heuristics
- Constraint propagation
- Smart variable selection

---

### 🟦 Piscine 12 — Latin Square Solver (7/10)

---

### 🟦 Piscine 13 — Kakuro Mini Solver (8/10)

---

### 🟩 PHASE 6 PROJECT — Irregular Sudoku Solver (9/10)

**Trick:**
- Regions are NOT standard 3×3 boxes

---

# 🟥 PHASE 7 — FINAL BOSS PROJECT (Difficulty: 10/10)

## TETRIS OPTIMIZER — Go Backtracking Masterpiece

### Goal
Place tetrominoes into the **smallest possible square grid** using optimized backtracking.

### Required Concepts
- Shape rotation & normalization
- Duplicate shape detection
- Heavy pruning
- Grid resizing
- Best-solution tracking

### Final Trick
Two different pieces may be **rotationally equivalent** — detect and skip duplicates.

---

## 🏁 Completion Outcome

After completing this Piscine, you will:

- Fully understand backtracking from fundamentals to optimization
- Be able to solve NP-style search problems
- Be ready for competitive programming, interviews, or advanced algorithm projects

---

## 🚀 Suggested Next Steps

- Implement test cases for each exercise
- Benchmark pruning effectiveness
- Convert Tetris Optimizer into a visual program
- Push everything into a GitHub portfolio

---

**Good luck — and welcome to Backtracking Hell 😄**

