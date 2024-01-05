// functions/functions.go
package functions

import (
	"fmt"
	"strconv"
	"unicode"
)

type Piece struct {
	name  string
	color string
	value int
}

func PrintBoard(board [8][8]Piece) {
	fmt.Println("\nBoard:")
	for _, row := range board {
		for _, piece := range row {
			if(piece.name == "None"){
				fmt.Print("  ")
			} else {
				fmt.Print(piece.name + " ")
			}

		}
		fmt.Println("")
	}
}
func GetColor(board [8][8]Piece, row, col int)int {
	if(board[row][col].color == "white"){
		return -1
	}
	return 1
}
func ConvertPGN(board [8][8]Piece, fen string) [8][8]Piece {
	values := map[string]int{
		"p": 1,
		"n": 3,
		"b": 3,
		"r": 5,
		"q": 9,
		"k": 10000,
		"P": 1,
		"N": 3,
		"B": 3,
		"R": 5,
		"Q": 9,
		"K": 10000,
	}

	var i, j int

	for _, character := range fen {
		char := string(character)

		if char == " "{
			break
		}
		if char == "/" || j == 8 {
			i++
			j = 0
		}

		if unicode.IsDigit(character) {
			stop, _ := strconv.Atoi(char)
			stop += j

			for j < stop && j < 8 {
				board[i][j] = Piece{name: "None", color: "None", value: 0}
				j++
			}

		} else if char != "/" {
			var pieceColor string
			if unicode.IsUpper(character) {
				pieceColor = "white"
			} else {
				pieceColor = "black"
			}
			board[i][j] = Piece{name: char, color: pieceColor, value: values[char]}
			j++
		}
	}

	return board
}
func ExampleBoard() string {
	return "r1bqkb1r/ppp2ppp/2n5/3np3/8/2NP1N2/PPP2PPP/R1BQKB1R"
}


//rn6/1p3k2/3p2p1/p4b1p/P1Pp4/8/1P3qPP/R2Q3K
