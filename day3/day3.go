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
	for scanner.Scan() {
		line := scanner.Text()
		needles, haystack := split_line(&line)
		match, err := common_letter(needles, haystack)
		if err != nil {
			log.Fatal(err)
		}

		charCodeAt := int(match[0])

		if charCodeAt < 97 {
			result += charCodeAt - 38
		} else {
			result += charCodeAt - 96
		}
	}

	fmt.Println(result)

}
