package main

import (
    "game/games"
    "game/ui"
)

func main() {
    // Initialize board and players
    board := games.NewBoard()
    player1 := games.NewPlayer("Player 1", "X")
    player2 := games.NewPlayer("Player 2", "O")
    
    // Start the game using CLI
    ui.StartCLI(board, player1, player2)
}
