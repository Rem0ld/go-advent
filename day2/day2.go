package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	win      = 6
	draw     = 3
	loss     = 0
	rock     = 1
	paper    = 2
	scissors = 3
)

type Rpc string

const (
	A Rpc = "A" // Rock
	B     = "B" // Paper
	C     = "C" // Scissors
	X     = "X" // Rock
	Y     = "Y" // Paper
	Z     = "Z" // Scissors
)

var game = map[Rpc]map[Rpc]int{
	A: { // Rock
		X: loss + scissors,
		Y: draw + rock,
		Z: win + paper,
	},
	B: { // Paper
		X: loss + rock,
		Y: draw + paper,
		Z: win + scissors,
	},
	C: { // Scissors
		X: loss + paper,
		Y: draw + scissors,
		Z: win + rock,
	},
}

type Round struct {
	total int
}

func New() *Round {
	r := Round{
		total: 0,
	}

	return &r
}

func (r *Round) compare(arr []string) {
	result := game[Rpc(arr[0])][Rpc(arr[1])]

	r.total += result
}

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	r := New()
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		if len(s) == 2 {
			r.compare(s)
		}
	}

	fmt.Println(r.total)
}
