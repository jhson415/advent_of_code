package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	one = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

func main() {
	var sum int
	numMap := map[string]int{
		"one":   one,
		"two":   two,
		"three": three,
		"four":  four,
		"five":  five,
		"six":   six,
		"seven": seven,
		"eight": eight,
		"nine":  nine,
	}
	re := regexp.MustCompile("[0-9]")

	s := readFile("puzzle_input.txt")
	ss := strings.Split(s, "\n")

	for i, v := range ss {
		fmt.Printf("[%d] string before conversion: %s\n", i, v)
		for s, n := range numMap {
			v = strings.ReplaceAll(v, s, s+strconv.Itoa(n)+s)
		}

		a := re.FindAllString(v, -1)

		n, err := strconv.Atoi(a[0] + a[len(a)-1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("[%d] found value: %d, input string: %s\n", i, n, v)
		sum = sum + n
	}
	fmt.Println(sum)

}

func readFile(Fpath string) string {
	file, err := os.Open(Fpath)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed to read file, %w", err))
	}
	return string(b)
}
