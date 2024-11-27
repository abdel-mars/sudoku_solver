# Sudoku Solver 🧩

A **Go-based Sudoku solver** that reads a Sudoku board, validates the input, and solves the puzzle using a **backtracking algorithm**.

![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)
![License](https://img.shields.io/badge/License-MIT-blue)

## 🚀 Features
- ✅ Validates Sudoku board input
- 🔄 Solves puzzles using the **backtracking algorithm**
- 🎮 Simple command-line interface for testing and usage
- 🧩 Easily adaptable for different Sudoku puzzles

## 📖 Table of Contents
- [Project Overview](#project-overview)
- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Contributing](#contributing)
- [License](#license)

## 🚀 Project Overview
This project implements a **Sudoku solver** using the **Go programming language** and solves puzzles using a **backtracking algorithm** while adhering to the basic rules of Sudoku:
- Each number (1-9) must appear only once in each row.
- Each number (1-9) must appear only once in each column.
- Each number (1-9) must appear only once in each 3x3 subgrid.

### Example Input:
The input for the Sudoku puzzle should be passed as strings. Empty cells should be represented as `0` or `.`. Here's an example of a valid Sudoku input:

```plaintext
53..7....
6..195...
.98....6.
8...6...3
4..8.3..1
7...2...6
.6....28.
...419..5
....8..79



---

## ⚙️ Installation

### Prerequisites
Before you begin, ensure that **Go** is installed on your machine. If Go isn't installed yet, you can download it from the official [Go Documentation](https://go.dev/doc/install). Follow the instructions based on your operating system.

### Steps

1. Clone the repository
Start by cloning the repository to your local machine. This will create a local copy of the project:

bash
Copy code
git clone https://github.com/abdel-mars/sudoku_solver.git
cd sudoku_solver
2. Install dependencies
The project uses Go modules to manage dependencies. To install all the required dependencies, run the following command in your terminal:

bash
Copy code
go mod tidy
This will download all the necessary modules and dependencies to your local project.

3. Build the project
Now, let's build the project to generate the executable binary. Run the following command:

bash
Copy code
make build
This command will compile your Go files and create the sudoku_solver binary. You will use this binary to run the Sudoku solver.

🧪 Usage
Now that you've installed the project, you can use the Sudoku solver to solve puzzles!

Running the Solver (Default Puzzle)
To run the solver with the default puzzle, use this command:

bash
Copy code
make run
This will run the solver with a pre-configured Sudoku puzzle, solving it and printing the result in your terminal.

Running with Custom Input
If you'd like to test a custom Sudoku puzzle, pass it as an argument to the solver in the following format:

bash
Copy code
./sudoku_solver "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
Each string represents one row of the puzzle:

Numbers (1-9) are filled cells.
Dots (.) or zeros (0) represent empty cells.
Example Output
When the solver successfully solves the puzzle, the result will be displayed in your terminal like this:

plaintext
Copy code
5 3 4 6 7 8 9 1 2
6 7 2 1 9 5 3 4 8
1 9 8 3 4 2 5 6 7
8 5 9 7 6 1 4 2 3
4 2 6 8 5 3 7 9 1
7 1 3 9 2 4 8 5 6
9 6 1 5 3 7 2 8 4
2 8 7 4 1 9 6 3 5
3 4 5 2 8 6 1 7 9
🧪 Running Tests
Running All Tests
To run all the tests for this project (including edge cases and puzzle validation checks), execute the following command:

bash
Copy code
make test
This will run the entire test suite, checking all the functionality and ensuring everything is working correctly.

Running Specific Tests
If you'd like to run a specific test, such as checking the validity of the Sudoku input, use the -run flag to specify the test name. For example:

bash
Copy code
go test -v sudoku_test.go -run TestValidSudokuInput
