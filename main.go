package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type (
	Row map[rune]int

	Table struct {
		size  int
		lines []Row
	}
)

func (t *Table) grow(n int) {
	for i := len(t.lines); n >= i; i++ {
		t.lines = append(t.lines, make(Row))
	}
}

func (t *Table) Search(s string) float64 {
	var (
		line    int     = 0
		counter float64 = 0
	)

	for _, r := range s {
		next := len(t.lines[line])
		if next > 1 || (next == 1 && line == 0) {
			counter += 1
		}
		line = t.lines[line][r]
	}

	return counter
}

func (t *Table) Add(s string) {
	line := 0
	for _, r := range s {
		if line >= len(t.lines) {
			t.grow(line)
		}
		_, ok := t.lines[line][r]
		if !ok {
			t.lines[line][r] = t.size + 1
			t.size += 1
		}
		line = t.lines[line][r]
	}
}

func (t *Table) Display() {
	fmt.Println()
	for line, row := range t.lines {
		fmt.Print(line, ": ")
		for r, o := range row {
			fmt.Print(string(r), "->", o, " | ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func NewTable() *Table {
	return &Table{
		size:  0,
		lines: make([]Row, 0),
	}
}

func Solve(input io.Reader) {
	reader := bufio.NewReader(input)

	for {
		text, err := getNextLine(reader)
		if err != nil {
			break
		}

		n, err := strconv.Atoi(text)
		if err != nil {
			break
		}

		table := NewTable()
		words := make([]string, n)

		for i := 0; n > i; i++ {
			text, err := getNextLine(reader)
			if err != nil {
				break
			}

			table.Add(text + "$")
			words[i] = text
		}

		// table.Display()

		var sum float64 = 0
		for _, word := range words {
			sum += table.Search(word)
		}

		fmt.Printf("%.2f\n", sum/float64(n))
	}
}

func getNextLine(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func main() {
	input := `4
hello
hell
heaven
goodbye
3
hi
he
h
7
structure
structures
ride
riders
stress
solstice
ridiculous
`

	Solve(strings.NewReader(input))
}
