// functions/moves.go
package functions

func MakeMove(board [8][8]Piece, row, col, targetRow, targetCol int) [8][8]Piece {
    // Create a copy of the original board
    newBoard := board

    // Make the move on the copy
    newBoard[targetRow][targetCol] = board[row][col]
    newBoard[row][col] = Piece{name: "None", color: "None", value: 0}

    // Return the modified copy
    return newBoard
}