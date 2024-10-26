package games

import "fmt"

// Board struct with a slice of slices to represent the cells
type Board struct {
    Cells [][]string
}

// NewBoard initializes a new game board with empty cells
func NewBoard() *Board {
    return &Board{
        Cells: [][]string{
            {" ", " ", " "},
            {" ", " ", " "},
            {" ", " ", " "},
        },
    }
}

// UpdateBoard updates the board at the specified row and column with the player's mark
func (b *Board) UpdateBoard(row, col int, mark string) bool {
    if b.Cells[row][col] == " " { // Check if the cell is empty
        b.Cells[row][col] = mark
        return true
    }
    return false
}

// DisplayBoard prints the current state of the board to the console
func (b *Board) DisplayBoard() {
    fmt.Println("Current Board:")
    for _, row := range b.Cells {
        fmt.Print("|")
        for _, cell := range row {
            fmt.Printf(" %s |", cell) // Using the Cells field
        }
        fmt.Println()
        fmt.Println("-----------") // Separator between rows
    }
}