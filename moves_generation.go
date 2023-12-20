// functions/moves.go
package functions

/*
generate_moves(board)
|
move 1 -> generate_moves(board_after_move1)
move 2 -> generate_moves(board_after_move2)
move 3 -> generate_moves(board_after_move3)

max_val(eval(first_generation)) - min_val(eval(second_generation))
*/

func WhiteMoves(board [8][8]Piece) []([8][8]Piece) {
	var boards []([8][8]Piece)
	var movescount [6]int
	for i, row := range board {
		for j, piece := range row {
			switch piece.name {
			case "P":
				pawnMoves := PawnMoves(board, i, j, -1)
				movescount[0] += len(pawnMoves)
				for _, move := range pawnMoves {
					boards = append(boards, move)
				}
			case "N":
				knightMoves := KnightMoves(board, i, j)
				movescount[1] += len(knightMoves)
				for _, move := range knightMoves {
					boards = append(boards, move)
				}
			case "B":
				bishopMoves := BishopMoves(board, i, j)
				movescount[2] += len(bishopMoves)
				for _, move := range bishopMoves {
					boards = append(boards, move)
				}
			case "R":
				rookMoves := RookMoves(board, i, j)
				movescount[3] += len(rookMoves)
				for _, move := range rookMoves {
					boards = append(boards, move)
				}
			case "Q":
				queenMoves := QueenMoves(board, i, j)
				movescount[4] += len(queenMoves)
				for _, move := range queenMoves {
					boards = append(boards, move)
				}
			case "K":
				kingMoves := KingMoves(board, i, j)
				movescount[5] += len(kingMoves)
				for _, move := range kingMoves {
					boards = append(boards, move)
				}
			}
			
		}
	}
	return boards
}

func BlackMoves(board [8][8]Piece) []([8][8]Piece) {
	var boards []([8][8]Piece)
	var movescount [6]int
	for i, row := range board {
		for j, piece := range row {
			switch piece.name {
			case "p":
				pawnMoves := PawnMoves(board, i, j,1)
				movescount[0] += len(pawnMoves)
				for _, move := range pawnMoves {
					boards = append(boards, move)
				}
			case "n":
				knightMoves := KnightMoves(board, i, j)
				movescount[1] += len(knightMoves)
				for _, move := range knightMoves {
					boards = append(boards, move)
				}
			case "b":
				bishopMoves := BishopMoves(board, i, j)
				movescount[2] += len(bishopMoves)
				for _, move := range bishopMoves {
					boards = append(boards, move)
				}
			case "r":
				rookMoves := RookMoves(board, i, j)
				movescount[3] += len(rookMoves)
				for _, move := range rookMoves {
					boards = append(boards, move)
				}
			case "q":
				queenMoves := QueenMoves(board, i, j)
				movescount[4] += len(queenMoves)
				for _, move := range queenMoves {
					boards = append(boards, move)
				}
			case "k":
				kingMoves := KingMoves(board, i, j)
				movescount[5] += len(kingMoves)
				for _, move := range kingMoves {
					boards = append(boards, move)
				}
			}
			
		}
	}
	return boards
}

func PawnMoves(board [8][8]Piece, row, col, rowdif int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	if row+rowdif >= 0 && row+rowdif <= 7 && board[row + rowdif][col].value == 0 {
		newBoard := MakeMove(board, row, col, row+rowdif, col)
		if !IsInCheck(newBoard, GetColor(board, row, col)){
			boards = append(boards, newBoard)
		}
	}
	supposed_row := 1
	if rowdif < 0 {
		supposed_row = 6
	}
	if row + 2*rowdif >= 0 && row + 2 * rowdif <= 7 && board[row + 2*rowdif][col].value == 0 && row == supposed_row && board[row+rowdif][col].value == 0{
		newBoard := MakeMove(board, row, col, row+ 2 * rowdif, col)
		if !IsInCheck(newBoard, GetColor(board, row, col)){
			boards = append(boards, newBoard)
		}
	}
	
	if row+rowdif >= 0 && row+rowdif <= 7 && col-1 >= 0 && board[row+rowdif][col-1].color != board[row][col].color && board[row+rowdif][col-1].value != 0 {
		// Create a copy of the board and make the move
		newBoard := MakeMove(board, row, col, row+rowdif, col-1)
		if !IsInCheck(newBoard, GetColor(board, row, col)){
			boards = append(boards, newBoard)
		}
	}
	if row+rowdif >= 0 && row+rowdif <= 7 && col+1 <= 7 && board[row+rowdif][col+1].color != board[row][col].color && board[row+rowdif][col+1].value != 0{
		// Create a copy of the board and make the move
		newBoard := MakeMove(board, row, col, row+rowdif, col+1)
		if !IsInCheck(newBoard, GetColor(board, row, col)){
			boards = append(boards, newBoard)
		}
	}

	return boards
}


func KnightMoves(board [8][8]Piece, row, col int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	knightMoves := [][]int{
		{-2, -1}, {-2, 1},
		{-1, -2}, {-1, 2},
		{1, -2}, {1, 2},
		{2, -1}, {2, 1},
	}

	for _, move := range knightMoves {
		newRow, newCol := row+move[0], col+move[1]

		if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
			if board[newRow][newCol].color == "" || board[newRow][newCol].color != board[row][col].color {
				// Create a copy of the board and make the move
				newBoard := MakeMove(board, row, col, newRow, newCol)
				if !IsInCheck(newBoard, GetColor(board, row, col)){
					boards = append(boards, newBoard)
				}
			}
		}
	}

	return boards
}

func BishopMoves(board [8][8]Piece, row, col int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	bishopMoves := [][]int{
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for _, move := range bishopMoves {
		newRow, newCol := row, col

		for {
			newRow += move[0]
			newCol += move[1]

			if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
				break
			}

			if board[newRow][newCol].color == "None" {
				// Create a copy of the board and make the move
				newBoard := MakeMove(board, row, col, newRow, newCol)
				if !IsInCheck(newBoard, GetColor(board, row, col)){
					boards = append(boards, newBoard)
				}
			} else {
				if board[newRow][newCol].color != board[row][col].color && board[newRow][newCol].name != "k" {
					// Create a copy of the board and make the move
					newBoard := MakeMove(board, row, col, newRow, newCol)
					if !IsInCheck(newBoard, GetColor(board, row, col)){
						boards = append(boards, newBoard)
					}
				}
				break
			}
		}
	}

	return boards
}
func RookMoves(board [8][8]Piece, row, col int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	rookMoves := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, move := range rookMoves {
		newRow, newCol := row, col

		for {
			newRow += move[0]
			newCol += move[1]

			if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
				break
			}

			// Allow rook to continue scanning even if it encounters a piece
			if board[newRow][newCol].color == "" || board[newRow][newCol].value == 0 {
				// Create a copy of the board and make the move
				newBoard := MakeMove(board, row, col, newRow, newCol)
				if !IsInCheck(newBoard, GetColor(board, row, col)){
					boards = append(boards, newBoard)
				}
			} else if board[newRow][newCol].color != board[row][col].color {
				// Capture opponent's piece, then break
				newBoard := MakeMove(board, row, col, newRow, newCol)
				if !IsInCheck(newBoard, GetColor(board, row, col)){
					boards = append(boards, newBoard)
				}
				break
			} else {
				break // Friendly piece blocking the path
			}
		}
	}

	return boards
}


func QueenMoves(board [8][8]Piece, row, col int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	rookMoves := RookMoves(board, row, col)
	bishopMoves := BishopMoves(board, row, col)

	for _, move := range rookMoves {
		boards = append(boards, move)
	}

	for _, move := range bishopMoves {
		boards = append(boards, move)
	}

	return boards
}

func KingMoves(board [8][8]Piece, row, col int) []([8][8]Piece) {
	var boards []([8][8]Piece)
	kingMoves := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, /*{0, 0},*/ {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Get the color of the king
	kingColor := board[row][col].color

	for _, move := range kingMoves {
		newRow, newCol := row+move[0], col+move[1]
		if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
			if board[newRow][newCol].name == "None" || board[newRow][newCol].color != kingColor {
	
				newBoard := MakeMove(board, row, col, newRow, newCol)

				if !IsInCheck(newBoard, GetColor(board, row, col)) {
					boards = append(boards, newBoard)
				}
				
			}
		}
	}

	return boards
}

// Helper function to find the position of the king on the board
func FindKing(board [8][8]Piece, color string) [2]int {
	var kingPos [2]int

	for i, row := range board {
		for j, piece := range row {
			if (piece.name == "K" && color == "white") || (piece.name == "k" && color == "black") {
				kingPos[0], kingPos[1] = i, j
				return kingPos
			}
		}
	}

	// Return an invalid position if the king is not found
	return [2]int{-1, -1}
}