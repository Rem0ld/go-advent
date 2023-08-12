package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	left  int
	right int
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

	for i := first; i <= second; i++ {
		count++
	}

	return Section{
		left:  first,
		right: second,
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

	switch true {
	case first.count < second.count:
		if first.left >= second.left && first.right <= second.right {
			fmt.Println(first, second)
			return 1
		}
	case second.count < first.count:
		if second.left >= first.left && second.right <= first.right {
			fmt.Println(first, second)
			return 1
		}
	}

	return 0
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
