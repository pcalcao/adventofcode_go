package main

import (
	"fmt"
	"log"
	"utils"
)

func executeComp(lines []int, position int) int {
	var op = lines[position]
	if op == 99 {
		return lines[0]
	}
	var operand1 = lines[position+1]
	var operand2 = lines[position+2]
	var dest = lines[position+3]
	if op == 1 {
		lines[dest] = lines[operand1] + lines[operand2]
		return executeComp(lines, position+4)
	}
	if op == 2 {
		lines[dest] = lines[operand1] * lines[operand2]
		return executeComp(lines, position+4)
	}

	return 0
}

func findParams(lines []int, noun int, verb int) int {
	lines2 := append([]int(nil), lines...)

	lines2[1] = noun
	lines2[2] = verb
	return executeComp(lines2, 0)
}

func main() {
	lines, err := utils.ReadLineAsListInt("inputs/day2_input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	//fmt.Println(findParams(lines, 12, 2))

	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			var res = findParams(lines, i, j)
			if res == 19690720 {
				fmt.Println(100*i + j)
			}
		}
	}

}
