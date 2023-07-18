package main

import "fmt"

// States the fields of the board can take.
type FieldState int

const (
	// Field is not taken by any player.
	None FieldState = iota
	// Field is taken by player x.
	X
	// Field is taken by player o.
	O
)

// Convert a field state into a string. Returns either " " (space), x, or o for
// states None, X and O respectively.
//
// Returns ? if the field state is unknown.
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

// Returns the string representation of a field state for usage in a board view.
// Assumes 3 character long field representations.
//
// If keyMapMode is set to true, it will return "[<fieldNumber>]" (e.g., "[3]")
// if the field is not taken and "   " otherwise.
// If keyMapMode is set to false, it will return " <field.String()> " (e.g.,
// " x ") if the field is taken and "   " otherwise.
func fieldString(field FieldState, fieldNumber int, keyMapMode bool) string {
	if keyMapMode && field == None {
		return fmt.Sprintf("[%d]", fieldNumber)
	} else if keyMapMode || field == None {
		return "   "
	}

	return fmt.Sprintf(" %v ", field)
}
