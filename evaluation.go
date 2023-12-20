package functions

import (
	//"strconv"
	//"fmt"
	"strings"
)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsInCheck(board [8][8]Piece, color int) bool {
	var row, col int
	//fmt.Print("Debugger: IsInCheck was called with color parameter: " + strconv.Itoa(color))
	//fmt.Println(" ")

	// Find the king's position
	for r, rn := range board {
		for c, cn := range rn {
			if (cn.name == "K" && color == -1) || (cn.name == "k" && color == 1) {
				row, col = r, c
			}
		}
	}

	//fmt.Println("Debugger: Current color found: " + strconv.Itoa(GetColor(board, row, col)))
	//fmt.Println("Debugger: King is at position: " + strconv.Itoa(row) + " " + strconv.Itoa(col))

	// Check for pawn attacks
	pawnAttacks := [][2]int{{row + color, col - 1}, {row + color, col + 1}}
	for _, attack := range pawnAttacks {
		attackRow, attackCol := attack[0], attack[1]
		if attackRow >= 0 && attackRow <= 7 && attackCol >= 0 && attackCol <= 7 {
			attackPiece := board[attackRow][attackCol]
			if attackPiece.name != "None" && strings.ToLower(attackPiece.name) == "p" && GetColor(board, attackRow, attackCol) == -color {
				//fmt.Println("Debugger: Pawn attacking at " + strconv.Itoa(attackRow) + " " + strconv.Itoa(attackCol))
				return true
			}
		}
	}

	// Check for bishop or queen attacks (diagonals)
	bishopAttacks := [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, attack := range bishopAttacks {
		i, j := attack[0], attack[1]
		for k := 1; k <= 7; k++ {
			newRow, newCol := row+k*i, col+k*j
			if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
				break // Out of bounds
			}
			attackPiece := board[newRow][newCol]
			if attackPiece.name != "None" && (strings.ToLower(attackPiece.name) == "b" || strings.ToLower(attackPiece.name) == "q") &&
				GetColor(board, newRow, newCol) == -color {
				//fmt.Println("Debugger: Check by queen / bishop")
				return true
			} else if attackPiece.name != "None" {
				break // A piece is blocking the path
			}
		}
	}

	// Check for rook or queen attacks (horizontals and verticals)
	rookAttacks := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, attack := range rookAttacks {
		i, j := attack[0], attack[1]
		for k := 1; k <= 7; k++ {
			newRow, newCol := row+k*i, col+k*j
			if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
				break // Out of bounds
			}
			attackPiece := board[newRow][newCol]
			if attackPiece.name != "None" && (strings.ToLower(attackPiece.name) == "r" || strings.ToLower(attackPiece.name) == "q") &&
				GetColor(board, newRow, newCol) == -color {
				//fmt.Println("Debugger: Check by rook / queen")
				return true
			} else if attackPiece.name != "None" {
				break // A piece is blocking the path
			}
		}
	}

	// Check for knight attacks
	knightMoves := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	for _, move := range knightMoves {
		newRow, newCol := row+move[0], col+move[1]
		if newRow >= 0 && newRow <= 7 && newCol >= 0 && newCol <= 7 {
			attackPiece := board[newRow][newCol]
			if attackPiece.name != "None" && strings.ToLower(attackPiece.name) == "n" && GetColor(board, newRow, newCol) == -color {
				//fmt.Println("Debugger: Knight attacking at " + strconv.Itoa(newRow) + " " + strconv.Itoa(newCol))
				return true
			}
		}
	}

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
	if IsInCheck(board, -1) {
		score -= 100
	}

	if len(BlackMoves(board)) == 0 {
		score = 99999
	}
	return score
}