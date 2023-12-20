// main.go
package main

import (
	"fmt"
	"l0v3/functions"
	"time"
)



func main() {
	board := [8][8]functions.Piece{}
	board = functions.ConvertPGN(board, functions.ExampleBoard())

	//functions.PrintBoard(board)
	//board = functions.MakeMove(board, 0,0,1,1)
	//fmt.Println(" ")
	//functions.PrintBoard(board)
	//fmt.Println(" ")

	
	startTime := time.Now()

	depth := 20	 // Change the depth as needed
	nb := [8][8]functions.Piece{}

	for i := 0; i < depth; i++ {
		whitemoves := functions.WhiteMoves(board)
		maxscore := -9999
		for _, gen := range whitemoves {
			whitescore := functions.EvaluateBoard(gen)

			blackmoves := functions.BlackMoves(gen)
			minscore := 9999
			for _, blk := range blackmoves {
				secondscore := functions.EvaluateBoard(blk) // eval white after black's move
				if secondscore < minscore {
					minscore = secondscore
				}
			}
			if whitescore+minscore > maxscore {
				maxscore = whitescore + minscore
				nb = gen
			}
		}
	}

	// Print the final board after all moves
	functions.PrintBoard(nb)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)

	fmt.Printf("Time taken: %s\n", elapsed)

}
