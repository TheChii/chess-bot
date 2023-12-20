package main

import (
	"fmt"
	"l0v3/functions"
	"time"

)

func Minimax(board [8][8]functions.Piece, depth, alpha, beta, color int) int {
    if depth == 0 {
        return functions.EvaluateBoard(board)
    }

    var moves []([8][8]functions.Piece)
    if color == 1 {
        moves = functions.WhiteMoves(board)
    } else {
        moves = functions.BlackMoves(board)
    }

    if len(moves) == 0 {
        // No legal moves, evaluate current board
        return functions.EvaluateBoard(board)
    }

    if color == 1 {
        maxScore := -9999
        for _, move := range moves {
            score := Minimax(move, depth-1, alpha, beta, -color)
            if score > maxScore {
                maxScore = score
            }
            alpha = max(alpha, score)
            if beta <= alpha {
                break // Beta cut-off
            }
        }
        return maxScore
    } else {
        minScore := 9999
        for _, move := range moves {
            score := Minimax(move, depth-1, alpha, beta, -color)
            if score < minScore {
                minScore = score
            }
            beta = min(beta, score)
            if beta <= alpha {
                break // Alpha cut-off
            }
        }
        return minScore
    }
}


func main() {

	board := [8][8]functions.Piece{}
	//functions.PrintBoard(board)
	//board = functions.MakeMove(board, 0,0,1,1)
	//fmt.Println(" ")
	//functions.PrintBoard(board)
	//fmt.Println(" ")

	
	for true {
		var fennot string
		fmt.Scan(&fennot)
		board = functions.ConvertPGN(board, fennot)
		
		
		startTime := time.Now()

		depth := 5
		var nb [8][8]functions.Piece
		maxScore := -9999
		for _, move := range functions.WhiteMoves(board) {
			score := Minimax(move, depth-1, -9999, 9999, -1)
			if score > maxScore {
				maxScore = score
				nb = move
			}
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
	
		functions.PrintBoard(nb)
		fmt.Printf("Time taken: %s\n", elapsed)
	
		
	}
	
	
	
	
}