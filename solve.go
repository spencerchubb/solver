package main

import (
	"fmt"
	"os"
	"solver"
	"strconv"
	"time"
)

const KEY_SCRAMBLE = "-scramble"
const KEY_MOVES = "-moves"
const KEY_SOLUTIONS = "-solutions"

const DEFAULT_SCRAMBLE = ""
const DEFAULT_MOVES = "UFDLRB"
const DEFAULT_SOLUTIONS = "1"

func main() {
	args := os.Args

	argMap := make(map[string]string)
	argMap[KEY_SCRAMBLE] = DEFAULT_SCRAMBLE
	argMap[KEY_MOVES] = DEFAULT_MOVES
	argMap[KEY_SOLUTIONS] = DEFAULT_SOLUTIONS

	key := ""
	for i := 1; i < len(args); i++ {
		if i%2 == 1 {
			// Odd index means it should be a key
			if args[i] != KEY_SCRAMBLE && args[i] != KEY_MOVES && args[i] != KEY_SOLUTIONS {
				fmt.Printf("Invalid key: %s", args[i])
				return
			}
			key = args[i]
		} else {
			argMap[key] = args[i]
			key = ""
		}
	}

	if key != "" {
		fmt.Printf("Expected an argument for: %s", key)
		return
	}

	facelets := solver.SolvedFacelets()
	solver.PerformAlgorithm(&facelets, argMap[KEY_SCRAMBLE])

	// var moveNames []string
	// for _, char := range argMap[KEY_MOVES] {
	// 	switch char {
	// 	case 'U':
	// 		moveNames = append(moveNames, "U1", "U2", "U3")
	// 	case 'F':
	// 		moveNames = append(moveNames, "F1", "F2", "F3")
	// 	case 'D':
	// 		moveNames = append(moveNames, "D1", "D2", "D3")
	// 	case 'B':
	// 		moveNames = append(moveNames, "B1", "B2", "B3")
	// 	case 'L':
	// 		moveNames = append(moveNames, "L1", "L2", "L3")
	// 	case 'R':
	// 		moveNames = append(moveNames, "R1", "R2", "R3")
	// 	default:
	// 		fmt.Printf("Invalid character: %c", char)
	// 		return
	// 	}
	// }
	// moves := solver.MoveSubset(moveNames)

	var moves []int
	for _, char := range argMap[KEY_MOVES] {
		switch char {
		case 'U':
			moves = append(moves, 0, 1, 2)
		case 'F':
			moves = append(moves, 3, 4, 5)
		case 'D':
			moves = append(moves, 6, 7, 8)
		case 'B':
			moves = append(moves, 9, 10, 11)
		case 'L':
			moves = append(moves, 12, 13, 14)
		case 'R':
			moves = append(moves, 15, 16, 17)
		default:
			fmt.Printf("Invalid character: %c", char)
			return
		}
	}

	maxSolutions, err := strconv.Atoi(argMap[KEY_SOLUTIONS])
	if err != nil {
		fmt.Printf("Invalid integer: %s", argMap[KEY_SOLUTIONS])
		return
	}

	startTime := time.Now()

	solver.Solve(facelets, moves, maxSolutions, true)

	endTime := time.Now()
	fmt.Printf("Time: %v", endTime.Sub(startTime))
}
