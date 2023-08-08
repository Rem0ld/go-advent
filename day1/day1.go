package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var mc []int
	curr := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			mc = append(mc, curr)
			curr = 0
		} else {
			curr += number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(mc)
	l := len(mc)

	fmt.Println(mc[l-1] + mc[l-2] + mc[l-3])
}
