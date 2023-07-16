package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

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
