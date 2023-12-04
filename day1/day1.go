package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		var numbers []string
		for _, char := range text {
			if unicode.IsNumber(char) {
				numbers = append(numbers, string(char))
			}
		}

		if len(numbers) > 0 {
			n, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
			total += n
		}
	}

	return total
}

func Part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var total int
	var needles = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		var numbers []string
		for i, char := range text {
			if unicode.IsNumber(char) {
				numbers = append(numbers, string(char))
			} else {
				for numIdx, val := range needles {
					if strings.HasPrefix(text[i:], val) {
						numbers = append(numbers, strconv.Itoa(numIdx+1))
					}
				}
			}
		}

		if len(numbers) > 0 {
			n, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
			total += n
		}
	}

	return total
}
