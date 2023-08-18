package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	const MARKER = 14
	for scanner.Scan() {
		temp := make([]string, 0, MARKER)
		line := scanner.Text()

		for i, l := range line {
			if idx := slices.Index(temp, string(l)); idx != -1 {
				_, temp = temp[0:idx], temp[idx+1:]

			}

			temp = append(temp, string(l))
			if len(temp) == MARKER {
				fmt.Printf("character at %v, %s\n", i+1, string(l))
				break
			}
		}
	}
}
