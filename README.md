# TicTacToe: An Implementation of CS50  Introduction to Artificial Intelligence with GO

## Overview

This project is a Tic-Tac-Toe game implemented in Go, incorporating principles from CS50’s Introduction to Artificial Intelligence. The game features an intelligent AI opponent using the Minimax algorithm, allowing players to engage in a challenging and strategic experience.

## File Structure

The project is organized as follows:

```bash
tic-tac-toe/
├── go.mod             # Go module file
├── main.go            # Entry point
├── games/
│   ├── board.go       # Handles the game board and state
│   ├── game.go        # Main game logic (checking for winner, game status)
│   └── player.go      # Manages player info and turns
└── ui/
    ├── cli.go         # CLI interface (to start, display board, take input)
    └── web.go         # Placeholder for future web interface
```

## Functions Overview

### CLI Functions

- **`StartCLI(board *games.Board, player1, player2 *games.Player)`**: 
  - Manages the game loop, prompting players for input and alternating turns.
  - Displays the current game board after each move.
  - Checks for game end conditions (win or tie).

### Game Logic Functions

- **`DisplayBoard()`**: 
  - Prints the current state of the Tic-Tac-Toe board to the console.
  
- **`UpdateBoard(row, col int, mark string)`**: 
  - Updates the board with the player's mark ('X' or 'O') at the specified location if the move is valid.
  
- **`CheckWinner(board *Board, moveCount int) string`**: 
  - Evaluates the board after each move to determine if there is a winner.

- **`IsBoardFull(board *Board) bool`**: 
  - Checks if the board is completely filled, indicating a tie.

### AI Functions

- **`Minimax(board *Board, depth int, isMaximizing bool) int`**:
  - The core function of the Minimax algorithm that evaluates all possible future moves recursively.
  - Returns the score of the board based on the terminal state (win, loss, tie).

- **`GetBestMove(board *Board) (int, int)`**: 
  - Determines the best move for the AI player by calling the `Minimax` function.
  
## Adversarial Search

### Introduction

Adversarial search is a type of algorithm used in scenarios where one agent (the player) faces another agent (the opponent) with opposing goals. In games like Tic-Tac-Toe, the objective of the AI is to win while the human player attempts to prevent it.

### Minimax Algorithm

The Minimax algorithm is a decision-making algorithm for two-player games. It operates under the following principles:

- **Utility Values**:
  - **-1**: Represents a loss for the maximizing player (AI).
  - **0**: Represents a tie.
  - **+1**: Represents a win for the maximizing player (AI).

- **Game Representation**:
  - **S₀**: The initial state (empty 3x3 board).
  - **Players(s)**: Determines which player's turn it is based on the current state.
  - **Actions(s)**: Returns all legal moves available in the current state.
  - **Result(s, a)**: Returns the new state after a move is made.
  - **Terminal(s)**: Checks if the game has ended.
  - **Utility(s)**: Returns the utility value for terminal states.

### How Minimax Works

1. **Recursion**:
   - The algorithm recursively explores all possible future game states up to terminal states, evaluating the outcomes based on current player turns.

2. **Alternating Strategy**:
   - The maximizing player tries to maximize the score, while the minimizing player tries to minimize it. This is done through a recursive simulation of the opponent's best possible moves.

3. **Decision Making**:
   - At each step, the maximizing player selects the action that leads to the maximum score based on the values generated from future possible states.

### Components of the Algorithm

- **S₀**: Initial state (an empty 3x3 board).
- **Players(s)**: Function that returns the current player's turn.
- **Actions(s)**: Function that returns all legal moves available in the current state.
- **Result(s, a)**: Function that returns a new state after applying action `a` to state `s`.
- **Terminal(s)**: Function that checks if the game has ended, returning true if a player has won or if it's a tie.
- **Utility(s)**: Function that returns the utility value of a terminal state.

## Conclusion

This Tic-Tac-Toe project serves as both a fun game and an educational implementation of concepts from artificial intelligence, showcasing how to apply algorithms in a practical setting.

## Future Improvements

- Develop a web interface for easier accessibility.
- Enhance AI strategies for more challenging gameplay.

Feel free to explore the code and contribute!