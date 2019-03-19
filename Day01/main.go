package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("input.txt")

	Part1(input)
	fmt.Println()
	Part2(input)
}

func loadInput(fileName string) []int {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s := strings.Split(string(raw), "\n")
	o := make([]int, len(s))

	for k, v := range s {
		o[k], _ = strconv.Atoi(v)
	}

	return o
}

func Part1(input []int) {
	sum := 0
	for _, v := range input {
		sum += v
	}

	fmt.Println("Part 1")
	fmt.Println("\tSum:", sum)
}

func Part2(input []int) {
	visited := make(map[int]bool)

	sum := 0
	firstDouble := 0
	boundary := 10000
	found := false

	for i := 0; !found && i < boundary; i++ {
		for _, v := range input {
			sum += v

			if _, isset := visited[sum]; isset {
				firstDouble = sum
				found = true
				break
			}

			visited[sum] = true
		}
	}

	if !found {
		fmt.Println("Error occured, no doubles.")
		os.Exit(1)
	}

	fmt.Println("Part 2:")
	fmt.Println("\tFirst Double:", firstDouble)
}
