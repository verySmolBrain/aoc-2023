package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func strip_chars(s string) string {
	replaceMap := map[string]string{
		"zero":  "zero0zero",
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for old, new := range replaceMap {
		s = strings.Replace(s, old, new, -1)
	}

	re := regexp.MustCompile("[^0-9]+")
	return re.ReplaceAllString(s, "")
}

func main() {
	fmt.Println("Hello, world.")

	file, err := os.ReadFile("input/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(file)
	lines := regexp.MustCompile("\n").Split(content, -1)

	var calibration_nums []int

	for _, line := range lines {
		digit_only_str := strip_chars(line)

		if len(digit_only_str) < 1 {
			continue
		}

		first_char := string(digit_only_str[0])
		last_char := string(digit_only_str[len(digit_only_str)-1])

		combined_num_str := strings.Join([]string{first_char, last_char}, "")

		combined_num, err := strconv.Atoi(combined_num_str)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		calibration_nums = append(calibration_nums, combined_num)

		fmt.Println(combined_num)
		fmt.Println(line)
	}

	// sum all the numbers
	sum := 0
	for _, num := range calibration_nums {
		sum += num
	}

	fmt.Println(sum)
}
