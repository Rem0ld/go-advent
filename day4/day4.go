package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Section struct {
	list  []int
	count int
}

type Sections struct {
	first  *Section
	second *Section
}

func toInteger(str []string) (int, int) {
	first, err := strconv.Atoi(str[0])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(str[1])
	if err != nil {
		log.Fatal(err)
	}

	return first, second
}

func NewSection(str string) Section {
	splitted := strings.Split(str, "-")
	first, second := toInteger(splitted)

	count := 0

	list := make([]int, 0)
	for i := first; i <= second; i++ {
		list = append(list, i)
		count++
	}

	return Section{
		list:  list,
		count: count,
	}
}

func NewSections(str []string) Sections {
	first := NewSection(str[0])
	second := NewSection(str[1])

	return Sections{
		first:  &first,
		second: &second,
	}
}

func (s *Sections) Compare() int {
	first := s.first
	second := s.second
	result := 0

	for _, val := range first.list {
		if slices.Contains(second.list, val) {
			result++
			break
		}
	}

	return result
}

func main() {
	file, err := os.Open("./file")
	if err != nil {
		log.Fatal(err)
	}

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sections := NewSections(strings.Split(line, ","))

		result += sections.Compare()
	}

	fmt.Println(result)

}
