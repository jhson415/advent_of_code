package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sum int

	file, err := os.Open("puzzle_input.txt")
	re := regexp.MustCompile("[0-9]")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed to read file, %w", err))
	}

	s := strings.Split(string(b), "\n")

	for _, v := range s {
		a := re.FindAllString(v, -1)
		n, err := strconv.Atoi(a[0] + a[len(a)-1])
		if err != nil {
			panic(err)
		}
		sum = sum + n
	}
	fmt.Println(sum)

}
