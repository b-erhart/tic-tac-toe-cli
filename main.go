package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func initTurn() FieldState {
	if rand.Float64() >= 0.5 {
		return X
	} else {
		return O
	}
}

func readSelect() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, errors.New("unable to read input")
	}

	field, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		return 0, errors.New("not a valid number")
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
			fmt.Printf("Error: invlid input (%s)\n", err.Error())
			fmt.Println("Please enter one of the available numbers")
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
