package main

import (
	"fmt"
	"strings"
)

// 3x3 playing board.
type Board struct {
	fields      [3][3]FieldState
	turnsPlayed int
}

// Return a string representation of the playing board. Each line is stored in
// a separate string in the returned slice.
//
// If keyMapMode is set to true, numbers will be printed for fields that are not
// yet taken and spaces otherwise.
// If keyMapMode is set to false, the corresponding field states will be printed
// for each field.
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

// String representation of the board. Shows the board state on the left and
// available fields and thier numbers for selection on the right.
func (board Board) String() string {
	var strBuilder strings.Builder
	boardStrings := board.stringRepresentation(false)
	keyStrings := board.stringRepresentation(true)

	for i := 0; i < len(boardStrings); i++ {
		line := fmt.Sprintf("%s     %s\n", boardStrings[i], keyStrings[i])
		strBuilder.WriteString(line)
	}

	return strBuilder.String()
}

// Set a field to a given state. Returns an error if there is no field for the
// given field number or if the field is already taken.
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

// Calculate whether the game is already finished. If it is, and the game is not
// tied, return the field state matching the winning player (x or o), otherwise
// None.
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
