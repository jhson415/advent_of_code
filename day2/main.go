package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic("failed to open input")
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic("failed to read input")
	}
	rows := strings.Split(string(b), "\n")
	for _, row := range rows {
		gs := strings.Split(row, ":")

		gID, err := strconv.Atoi(strings.Split(gs[0], " ")[1])
		if err != nil {
			panic("failed to convert game ID")
		}
		fmt.Printf("Game ID: %d\n", gID)

		gsl := strings.Split(gs[1], ";")
		for _, g := range gsl {
			var cCount struct {
				red   int
				green int
				blue  int
			}
			var cs []string
			if strings.Contains(g, ",") {
				cs = strings.Split(g, ", ")
			} else {
				cs = append(cs, g)
			}
			fmt.Println(g)
			for _, c := range cs {
				// fmt.Println(cCount)
				c = strings.TrimSpace(c)
				f := strings.Split(c, " ")
				cNum, err := strconv.Atoi(f[0])
				if err != nil {
					panic("failed to convert number")
				}
				cCol := f[1]

				switch {
				case cCol == "red":
					cCount.red = cCount.red + cNum
				case cCol == "green":
					cCount.green = cCount.green + cNum
				case cCol == "blue":
					cCount.blue = cCount.blue + cNum
				}

				switch {
				case cCount.red > 12 || cCount.green > 13 || cCount.blue > 14:
					result = result + gID
					fmt.Printf("Failed! cCount: %+v, total: %d, gID:%d\n", cCount, result, gID)
					goto exit_loop
				}
			}
		}
	exit_loop:
	}

	fmt.Printf("final result: %d\n", result)

}
