package games



// CheckWinner checks the board for a winner
func CheckWinner(b *Board) string {
    // Check rows and columns for a win
    for i := 0; i < 3; i++ {
        if b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][1] == b.Cells[i][2] && b.Cells[i][0] != " " {
            return b.Cells[i][0] // Return the winner's mark
        }
        if b.Cells[0][i] == b.Cells[1][i] && b.Cells[1][i] == b.Cells[2][i] && b.Cells[0][i] != " " {
            return b.Cells[0][i] // Return the winner's mark
        }
    }
    // Check diagonals for a win
    if b.Cells[0][0] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][2] && b.Cells[0][0] != " " {
        return b.Cells[0][0]
    }
    if b.Cells[0][2] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][0] && b.Cells[0][2] != " " {
        return b.Cells[0][2]
    }
    return "" // No winner
}

// IsBoardFull checks if the board is full
func IsBoardFull(b *Board) bool {
    for _, row := range b.Cells {
        for _, cell := range row {
            if cell == " " {
                return false // Found an empty cell
            }
        }
    }
    return true // No empty cells found
}