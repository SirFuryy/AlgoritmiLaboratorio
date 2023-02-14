package main

import (
	"fmt"
)

func pressSwitch(lights []bool, switchIndex int) {
	if switchIndex==1 {
		lights[len(lights)-1]= !lights[len(lights)-1]
		lights[0] = !lights[0]
		lights[1] = !lights[1]
	} else if switchIndex==len(lights) {
		lights[switchIndex-2]= !lights[switchIndex-2]
		lights[switchIndex-1]= !lights[switchIndex-1]
		lights[0] = !lights[0]
	} else {
		lights[switchIndex-2] = !lights[switchIndex-2]
		lights[switchIndex-1] = !lights[switchIndex-1]
		lights[switchIndex] = !lights[switchIndex]
	}
}

func pressSequence(lights []bool, switches []int) {
	for _, switchIndex := range switches {
		pressSwitch(lights, switchIndex)
	}
}

func findSolution(lights []bool) []int {
	solution := []int{}
	for i := range lights {
		if lights[i] == true {
			pressSwitch(lights, i+1)
			solution = append(solution, i+1)
			fmt.Println(lights)
		}
	}
	return solution
}

func main() {
	lights := []bool{true, true, true, false, true, true, true}
	fmt.Println("Initial state:", lights)

	// Press switch 3
	pressSwitch(lights, 3)
	fmt.Println("After pressing switch 3:", lights)

	// Press switch sequence [1, 2, 4, 6]
	pressSequence(lights, []int{1, 2, 4, 6})
	fmt.Println("After pressing switch sequence [1, 2, 4, 6]:", lights)

	// Find solution to turn off all lights
	solution := findSolution(lights)
	fmt.Println("Solution to turn off all lights:", solution)

	// Minimum number of switches to turn off all lights
	fmt.Println("Minimum number of switches to turn off all lights:", len(solution))
}
