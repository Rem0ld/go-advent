package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func parseInstructions(line string) (quantity, from, to int) {
	re := regexp.MustCompile(`\d+`)
	splitted := re.FindAllString(line, -1)

	quantity, _ = strconv.Atoi(splitted[0])
	from, _ = strconv.Atoi(splitted[1])
	to, _ = strconv.Atoi(splitted[2])

	from--
	to--

	return
}

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
			fmt.Println(line)
		} else {
			instructions = append(instructions, line)
		}
	}

	lastEl, drawing := drawing[len(drawing)-1], drawing[:len(drawing)-1]
	numCol := len(strings.Split(lastEl, " ")) / 3

	// Creating empty columns
	var result [][]string
	for i := 0; i < numCol; i++ {
		result = append(result, make([]string, 0))
	}

	// filling columns
	re := regexp.MustCompile(`\w`)
	for i := 0; i < len(drawing); i++ {
		for j := 0; j < len(drawing[i]); j++ {
			if re.MatchString(string(drawing[i][j])) {
				letter := string(drawing[i][j])
				col := int(math.Ceil(float64(j / 4)))
				result[col] = append(result[col], letter)
			}
		}
	}

	for _, col := range result {
		slices.Reverse(col)
	}

	// reading instructions
	for _, line := range instructions {
		quantity, from, to := parseInstructions(line)
		fmt.Println(result)

		fromCol := result[from]
		temp := make([]string, 0)

		for i := 0; i < quantity; i++ {
			elems, col := fromCol[len(fromCol)-1], fromCol[:len(fromCol)-1]
			fromCol = col
			temp = append(temp, elems)
		}

		result[from] = fromCol
		slices.Reverse(temp)
		result[to] = append(result[to], temp...)
	}

	var str []string
	for _, line := range result {
		str = append(str, line[len(line)-1])
	}
	fmt.Println(strings.Join(str, ""))
}
