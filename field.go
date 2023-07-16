package main

import "fmt"

type FieldState int

const (
	None FieldState = iota
	X
	O
)

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
