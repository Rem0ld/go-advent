package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()
		count++
		if count == 3 {
			// look for badge

			count = 0
			lines = nil
		}
	}

	fmt.Println(result)
}
