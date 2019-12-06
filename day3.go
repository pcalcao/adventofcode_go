package main

import (
    "fmt"
    "log"
	"strconv"
	"strings"
	"adventofcode2019/utils"
	"math"
)

type coord struct {
    x int
    y int
}


func getPathCircuit(path string) map[coord]int {
	var instructions = strings.Split(path, ",")
	x := map[coord]int{}
	var position = coord{x: 0, y: 0}
	var steps = 0
	for _, instruction := range instructions {
		direction := string(instruction[0])
		moves, _ := strconv.Atoi(instruction[1:])
		moveList, newPosition := getMoveList(direction, moves, position)
		position = newPosition
		for _, move := range moveList {
			steps++
			_, seenBefore := x[move]
			if !seenBefore {
				x[move] = steps
			}
		}
	}
	return x
}

func getMoveList(direction string, steps int, latest coord) ([]coord, coord) {
	var moves []coord
	switch direction {
	case "R":
		// move right
		for i := 0; i < steps; i++ {
			latest = coord{x: latest.x + 1, y: latest.y}
			moves = append(moves, latest)
		}
	case "L":
		// move left
		for i := 0; i < steps; i++ {
			latest = coord{x: latest.x - 1, y: latest.y}
			moves = append(moves, latest)
		}
	case "U":
		// move right
		for i := 0; i < steps; i++ {
			latest = coord{x: latest.x, y: latest.y + 1}
			moves = append(moves, latest)
		}
	case "D":
		// move right
		for i := 0; i < steps; i++ {
			latest = coord{x: latest.x, y: latest.y - 1}
			moves = append(moves, latest)
		}
	}
	return moves, latest
}

func intersection(c1 map[coord]int, c2 map[coord]int) []coord {
	var inter []coord
	for key := range c2 {
		_, exists := c1[key]
		if (exists) {
			inter = append(inter, key)
		}
	}
	return inter
}

// Abs : Go is stupid and doesn't have Abs for ints
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInputs() []string {
	lines, err := utils.ReadLines("inputs/day3_input")
	if err != nil {
        log.Fatalf("readLines: %s", err)
	}

	return lines
}

func minDistance(points []coord) int {
	minVal := math.MaxInt32
	for _, coordinate := range points {
		distance := Abs(coordinate.x) + Abs(coordinate.y)
		if (distance < minVal) {
			minVal = distance 
		}
	}
	return minVal
}

func minCombinedDistance(points []coord, c1 map[coord]int, c2 map[coord]int) int {
	minVal := math.MaxInt32
	for _, coordinate := range points {
		distance := c1[coordinate] + c2[coordinate]
		if (distance < minVal) {
			minVal = distance 
		}
	}
	return minVal
}

func day1Part1(c1 string, c2 string) {
	path1 := getPathCircuit(c1)
	path2 := getPathCircuit(c2)
	intersect := intersection(path1, path2)
	fmt.Println(minDistance(intersect))
}

func day1Part2(c1 string, c2 string) {
	path1 := getPathCircuit(c1)
	path2 := getPathCircuit(c2)
	intersect := intersection(path1, path2)
	fmt.Println(minCombinedDistance(intersect, path1, path2))
}

func main() {
	lines := getInputs()
	day1Part1(lines[0], lines[1])
	day1Part2(lines[0], lines[1])

	
}