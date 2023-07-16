package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func initTurn() FieldState {
	if rand.Float64() >= 0.5 {
		return X
	} else {
		return O
	}
}

func readSelect() (int, error) {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil && err.Error() != "unexpected newline" && err.Error() != "expected newline" {
		fmt.Printf("Unable to read input, exiting. (%v)\n", err)
		os.Exit(1)
	}

	if err != nil {
		return 0, errors.New("invalid input")
	}

	field, err := strconv.Atoi(input)

	if err != nil {
		return 0, errors.New("invalid input")
	}

	return field, nil
}

func printGameOver(winner FieldState, board Board) {
	if winner != None {
		fmt.Println("\n=== GAME OVER ===")
		fmt.Print(board)
		fmt.Printf("Player %v wins!", winner)
	} else {
		fmt.Println("\n=== GAME OVER ===")
		fmt.Print(board)
		fmt.Printf("Tie!")
	}
}

func main() {
	fmt.Print("Welcome to Tic Tac Toe!\n\n")

	var board Board
	turn := initTurn()

	for {
		fmt.Print(board)
		fmt.Println()

		fmt.Printf("It's player %v's turn!\n", turn)
		fmt.Print("Choose a field number: ")

		selectedField, err := readSelect()

		if err != nil {
			fmt.Println("Invlid input. Please enter one number.")
			continue
		}

		err = board.Set(selectedField, turn)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		finished, winner := board.GameFinished()

		if finished {
			printGameOver(winner, board)
			break
		}

		if turn == X {
			turn = O
		} else {
			turn = X
		}
	}
}
