package main

import (
	"fmt"
	"log"
	"strconv"
	"utils"
)

func costOfFuel(weight int) int {
	return (weight/3 - 2)
}

func costOfFuelWithFuel(weight int, total int) int {
	cost := costOfFuel(weight)
	if cost <= 0 {
		return total
	}
	return costOfFuelWithFuel(cost, total+cost)
}

func fuelForModules(modules []string) int {
	var sum int
	for _, module := range modules {
		weight, _ := strconv.Atoi(module)
		sum += costOfFuel(weight)
	}
	return sum
}

func fuelForModules2(modules []string) int {
	var sum int
	for _, module := range modules {
		weight, _ := strconv.Atoi(module)
		sum += costOfFuelWithFuel(weight, 0)
	}
	return sum
}

func main() {
	lines, err := utils.ReadInput(2019, 1)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(fuelForModules(lines))
	fmt.Println(fuelForModules2(lines))
}
