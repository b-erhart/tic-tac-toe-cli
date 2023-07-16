package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type FieldState int

const (
	None FieldState = iota
	X
	O
)

type Board struct {
	fields [3][3]FieldState
	turnsPlayed int
}

func (field FieldState) String() string {
	switch field {
	case None:
		return " "
	case X:
		return "x"
	case O:
		return "o"
	}

	return "?"
}

func fieldString(field FieldState, fieldId int, keyMapMode bool) string {
	if keyMapMode && field == None {
		return fmt.Sprintf("[%d]", fieldId)
	} else if keyMapMode || field == None {
		return "   "
	}

	return fmt.Sprintf(" %v ", field)
}

func newBoard() Board {
	return Board{[3][3]FieldState{}, 0}
}

func (board Board) stringRepresentation(keyMapMode bool) []string {
	lines := make([]string, len(board.fields)*2-1)

	for i, row := range board.fields {
		var strBuilder strings.Builder

		for j, field := range row {
			strBuilder.WriteString(fieldString(field, i*3+j, keyMapMode))

			if j+1 < len(row) {
				strBuilder.WriteRune('|')
			}
		}

		lines[i*2] = strBuilder.String()

		if i+1 < len(board.fields) {
			lines[i*2+1] = "---+---+---"
		}
	}

	return lines
}

func (board Board) String() string {
	var strBuilder strings.Builder
	boardStrings := board.stringRepresentation(false)
	keyStrings := board.stringRepresentation(true)

	for i := 0; i < len(boardStrings); i++ {
		strBuilder.WriteString(fmt.Sprintf("%s     %s\n", boardStrings[i], keyStrings[i]))
	}

	return strBuilder.String()
}

func (board *Board) Set(fieldNumber int, state FieldState) error {
	if fieldNumber < 0 || fieldNumber > 9 {
		return fmt.Errorf("%d is not a field number", fieldNumber)
	}

	i := fieldNumber / 3
	j := fieldNumber % 3

	if board.fields[i][j] != None {
		return fmt.Errorf("field %d is already taken", fieldNumber)
	}

	board.fields[i][j] = state
	board.turnsPlayed++
	return nil
}

func (board Board) GameFinished() (bool, FieldState) {
RowLoop:
	for _, row := range board.fields {
		rowState := row[0]

		for _, field := range row {
			if field == None || field != rowState {
				continue RowLoop
			}
		}

		return true, rowState
	}

ColLoop:
	for j := 0; j < 3; j++ {
		colState := board.fields[0][j]

		for i := 0; i < 3; i++ {
			if board.fields[i][j] == None || board.fields[i][j] != colState {
				continue ColLoop
			}
		}

		return true, colState
	}

	if board.fields[0][0] == board.fields[1][1] && board.fields[1][1] == board.fields[2][2] && board.fields[0][0] != None {
		return true, board.fields[0][0]
	}

	if board.fields[0][2] == board.fields[1][1] && board.fields[1][1] == board.fields[2][0] && board.fields[0][2] != None {
		return true, board.fields[0][2]
	}

	if board.turnsPlayed >= 9 {
		return true, None
	}

	return false, None
}

func main() {
	fmt.Print("Welcome to Tic Tac Toe!\n\n")

	var board Board
	turn := X

	if rand.Float64() >= 0.5 {
		turn = O
	}

	for {
		fmt.Print(board)
		fmt.Println()

		fmt.Printf("It's player %v's turn!\n", turn)
		fmt.Print("Choose a field number: ")

		var input string
		_, err := fmt.Scanln(&input)

		if err != nil && err.Error() != "unexpected newline" && err.Error() != "expected newline" {
			fmt.Printf("Unable to read input, exiting. (%v)\n", err)
			os.Exit(1)
		}

		if err != nil {
			fmt.Println("Invalid input. Please enter one number.")
			continue
		}

		selectedField, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter one number.")
			continue
		}

		err = board.Set(selectedField, turn)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		finished, winner := board.GameFinished()

		if finished && winner != None {
			fmt.Println("\n=== GAME OVER ===")
			fmt.Print(board)
			fmt.Printf("Player %v wins!", winner)
			break
		} else if finished {
			fmt.Println("\n=== GAME OVER ===")
			fmt.Print(board)
			fmt.Printf("Tie!")
			break
		}

		if turn == X {
			turn = O
		} else {
			turn = X
		}
	}
}
