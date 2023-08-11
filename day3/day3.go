package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func split_line(line *string) ([]string, string) {
	half := len(*line) / 2
	splitted := strings.Split(*line, "")
	x := make([]string, half)
	copy(x, splitted[0:half])

	return x, strings.Join(splitted[half:], "")
}

func common_letter(needles []string, haystack string) (string, error) {
	for _, needle := range needles {
		if strings.Contains(haystack, needle) {
			return needle, nil
		}
	}

	return "", errors.New("No match")
}

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var result int
	count := 0
	var lines []string
	var longest []string

	for scanner.Scan() {
		line := scanner.Text()
		if count == 0 {
			fmt.Println("longest should be nil", longest)
		}

		if len(line) > len(longest) {
			lines = append(lines, strings.Join(longest, ""))
			longest = strings.Split(line, "")
		} else {
			lines = append(lines, line)
		}

		count++

		if count == 3 {
			slices.Sort(longest)
			longest = slices.Compact(longest)
			fmt.Println(longest)

			for _, val := range longest {
				fmt.Println(val, int(val[0]))
				if strings.Contains(lines[0], val) && strings.Contains(lines[1], val) {
					fmt.Println(val, int(val[0]))
					result += int(val[0])
					break
				}
			}
			count = 0
			lines = nil
			longest = nil
		}
	}

	fmt.Println(result)
}
