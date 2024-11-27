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
1. **Clone the repository**:
   Start by cloning the repository to your local machine:
   ```bash
   git clone https://github.com/abdel-mars/sudoku_solver.git
   cd sudoku_solver
