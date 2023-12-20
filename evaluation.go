package functions

func IsInCheck(board [8][8]Piece, color int) bool {
	var row, col int

	// Find the king's position
	for r, rn := range board {
		for c, cn := range rn {
			if (cn.name == "K" && color == 1) || (cn.name == "k" && color == -1) {
				row, col = r, c
			}
		}
	}

	// Check for pawn attacks
	if row+color >= 0 && row+color <= 7 {
		if col-1 >= 0 && board[row+color][col-1].color != board[row][col].color && board[row+color][col-1].name != "" && (board[row+color][col-1].name == "P" || board[row+color][col-1].name == "p") {
			return true
		}
		if col+1 <= 7 && board[row+color][col+1].color != board[row][col].color && board[row+color][col+1].name != "" && (board[row+color][col+1].name == "P" || board[row+color][col+1].name == "p") {
			return true
		}
	}

	// Check for rook attacks
	for i := -1; i <= 1; i += 2 { // Iterate over left and right (-1 and 1)
		for j := -1; j <= 1; j += 2 { // Iterate over up and down (-1 and 1)
			for k := 1; k <= 7; k++ {
				newRow, newCol := row+k*i, col+k*j
				if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
					break // Out of bounds
				}
				if board[newRow][newCol] != (Piece{}) {
					if (board[newRow][newCol].name == "R" || board[newRow][newCol].name == "r" || board[newRow][newCol].name == "Q" || board[newRow][newCol].name == "q") &&
						board[newRow][newCol].color != board[row][col].color {
						return true
					} else {
						break // A piece is blocking the path
					}
				}
			}
		}
	}

	// Check up-left
	for i := -1; i <= 1; i += 2 { // Iterate over left and right (-1 and 1)
		for j := -1; j <= 1; j += 2 { // Iterate over up and down (-1 and 1)
			for k := 1; k <= 7; k++ {
				newRow, newCol := row+k*i, col+k*j
				if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
					break // Out of bounds
				}
				if board[newRow][newCol] != (Piece{}) {
					if (board[newRow][newCol].name == "R" || board[newRow][newCol].name == "r" || board[newRow][newCol].name == "Q" || board[newRow][newCol].name == "q") &&
						board[newRow][newCol].color != board[row][col].color {
						return true
					} else {
						break // A piece is blocking the path
					}
				}
			}
		}
	}

	// Check for bishop or queen attacks (diagonals)
	for i := -1; i <= 1; i += 2 { // Iterate over left and right (-1 and 1)
		for j := -1; j <= 1; j += 2 { // Iterate over up and down (-1 and 1)
			for k := 1; k <= 7; k++ {
				newRow, newCol := row+k*i, col+k*j
				if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
					break // Out of bounds
				}
				if board[newRow][newCol] != (Piece{}) {
					if (board[newRow][newCol].name == "B" || board[newRow][newCol].name == "b" || board[newRow][newCol].name == "Q" || board[newRow][newCol].name == "q") &&
						board[newRow][newCol].color != board[row][col].color {
						return true
					} else {
						break // A piece is blocking the path
					}
				}
			}
		}
	}
	// If no checks were found, return false
	return false
}


func EvaluateBoard(board [8][8]Piece) int{
	/*
	f(p) =
       + 9(Q-Q')
       + 5(R-R')
       + 3(B-B' + N-N')
       + 1(P-P')
       - 0.5(D-D' + S-S' + I-I')
       + 0.1(M-M') + ...

		KQRBNP = number of kings, queens, rooks, bishops, knights and pawns
		D,S,I = doubled, blocked and isolated pawns
		M = Mobility (the number of legal moves)
	
	var score int
	*/
	var score int
	var piecesArr = [6]int{0,0,0,0,0,0}

	//if IsInCheck(board, -1) {
	//	return -10000 
	//} else if IsInCheck(board, 1) {
	//	score += 1 
	//}
	
	for _, row := range board {
		for _, piece := range row {
			switch piece.name {
			case "Q":
				piecesArr[0]++ 
			case "q":
				piecesArr[0]-- 
			case "R":
				piecesArr[1]++
			case "r":
				piecesArr[1]--
			case "B":
				piecesArr[2]++ 
			case "b":
				piecesArr[2]--
			case "N":
				piecesArr[3]++ 
			case "n":
				piecesArr[3]-- 
			case "P":
				piecesArr[4]++ 
			case "p":
				piecesArr[4]-- 
			}
		}
	}

	var d, D, s, S int
	for index, row := range board{
		for col, piece := range row {
			if piece.name == "p"{
				if board[index-1][col].name == "p" {
					d++
				}
				if index + 1 <= 7 && board[index + 1][col].color != "None"{
					s++
				}

			} else if piece.name == "P"{
				if board[index+1][col].name == "P" {
					D++
				}
				if index - 1 >= 0 && board[index - 1][col].color != "None"{
					S++
				}
			}
		}
	}

	score = piecesArr[4] + 3 * (piecesArr[2] + piecesArr[3]) + 5 * piecesArr[1] + 9 * piecesArr[0]
	return score
}