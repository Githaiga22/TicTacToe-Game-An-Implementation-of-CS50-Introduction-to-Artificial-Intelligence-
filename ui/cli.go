package ui

import (
    "fmt"
    "game/games"
)

func StartCLI(board *games.Board, player1, player2 *games.Player) {
    current := player1
    moveCount := 0 // Track the number of moves
    for {
        board.DisplayBoard() // This now works since DisplayBoard is a method of Board
        
        // Prompt for player move
        var row, col int
        fmt.Printf("%s's turn (%s): Enter row and column (0, 1, or 2): ", current.Name, current.Mark)

        // Validate input
        if _, err := fmt.Scanf("%d %d", &row, &col); err != nil {
            fmt.Println("Invalid input! Please enter two integers.")
            fmt.Scanf("%*s") // Clear the invalid input
            continue
        }

        // Validate move
        if row < 0 || row > 2 || col < 0 || col > 2 || !board.UpdateBoard(row, col, current.Mark) {
            fmt.Println("Invalid move! Try again.")
            continue
        }

        moveCount++ // Increment move count after a valid move

        // Check for a winner
        if winner := games.CheckWinner(board); winner != "" {
            board.DisplayBoard() // This will now display the final board
            fmt.Printf("Congratulations %s! You win!\n", current.Name)
            break
        }

        // Check for a full board (tie)
        if games.IsBoardFull(board) {
            board.DisplayBoard()
            fmt.Println("It's a tie!")
            break
        }

        // Switch turns
        if current == player1 {
            current = player2
        } else {
            current = player1
        }
    }
}


// displayBoard prints the current state of the board to the console
// displayBoard prints the current state of the board to the console
// Assuming your Board struct looks something like this
type Board struct {
    Cells [3][3]string // Example field for a 3x3 board
}

// DisplayBoard prints the current state of the board to the console
func (b *Board) DisplayBoard() {
    fmt.Println("Current Board:")
    for i := 0; i < 3; i++ {
        fmt.Print("|")
        for j := 0; j < 3; j++ {
            fmt.Printf(" %s |", b.Cells[i][j]) // Using the Cells field
        }
        fmt.Println()
        fmt.Println("-----------") // Separator between rows
    }
}
