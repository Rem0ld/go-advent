package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}

	var drawing []string
	var instructions []string
	mode := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			mode = 1
			continue
		}

		if mode == 0 {
			drawing = append(drawing, line)
		} else {
			instructions = append(instructions, line)
		}
	}

	lastEl, drawing := drawing[len(drawing)-1], drawing[:len(drawing)-1]
	numCol := len(strings.Split(lastEl, " ")) / 3
	fmt.Println(drawing)
	fmt.Println(numCol)

	// We have everything to start filling columns
	// number of columns and drawing in a slice
	// we still need to split each line
	// we know it will come as a multiple of 3
	// 2 solutions that comes to my mind
	// we use something like split (strings or slices method) this should output [3]string
	// or we loop by len(slice) / 3, this should give us column by column
	// we need two loop i,j.
}
