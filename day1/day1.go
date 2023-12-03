package day1

import (
	"strconv"
	"strings"

	"github.com/larsha/aoc-2023/v2/utils"
)

var needles = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Part1() int {
	file := utils.ReadFile()

	var total int
	for _, str := range file {
		var positions = make([]string, len(str))

		for _, needle := range needles {
			firstIndex := strings.Index(str, needle)
			if firstIndex >= 0 {
				positions[firstIndex] = needle
			}

			lastIndex := strings.LastIndex(str, needle)
			if lastIndex >= 0 {
				positions[lastIndex] = needle
			}
		}

		var numbers []string
		for _, p := range positions {
			if p != "" {
				numbers = append(numbers, p)
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
	extraNeedles := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	file := utils.ReadFile()

	var total int
	for _, str := range file {
		var positions = make([]string, len(str))

		for i, needle := range extraNeedles {
			index := strings.Index(str, needle)
			if index >= 0 {
				positions[index] = needles[i]
			}

			lastIndex := strings.LastIndex(str, needle)
			if lastIndex >= 0 {
				positions[lastIndex] = needles[i]
			}
		}

		for _, needle := range needles {
			index := strings.Index(str, needle)
			if index >= 0 {
				positions[index] = needle
			}

			lastIndex := strings.LastIndex(str, needle)
			if lastIndex >= 0 {
				positions[lastIndex] = needle
			}
		}

		var numbers []string
		for _, char := range positions {
			if char != "" {
				numbers = append(numbers, char)
			}
		}

		n, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		total += n
	}

	return total
}
