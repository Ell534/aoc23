package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var unicodeToNumber = map[byte]int{
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}

var wordToNumber = map[string]string{
	"zero":  "z0ero",
	"one":   "o1ne",
	"two":   "t2wo",
	"three": "t3hree",
	"four":  "f4our",
	"five":  "f5ive",
	"six":   "s6ix",
	"seven": "s7even",
	"eight": "e8ight",
	"nine":  "n9ine",
}

// take the input
// split into each line
// extract 'digit' from each line
// combine first and last digit from each line into one number i.e 3 and 6 becomes 36
// sum all combined digits

func fileReader(pathToFile string) []string {

	var scanSlice = []string{}
	file, err := os.Open(pathToFile)

	if err != nil {
		fmt.Println(errors.New("File not found"))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanSlice = append(scanSlice, scanner.Text())
	}
	return scanSlice
}

func wordToDigit(stringSlice []string) []string {
	for key, value := range wordToNumber {
		for i, string := range stringSlice {
			stringSlice[i] = strings.ReplaceAll(string, key, value)
		}
	}
	return stringSlice
}

func extractDigits(stringSlice []string) []string {
	stringDigits := make([]string, len(stringSlice))
	for stringSliceIndex, word := range stringSlice {
		for _, char := range word {
			if unicode.IsDigit(char) {
				stringDigits[stringSliceIndex] = stringDigits[stringSliceIndex] + string(char)
			}
		}
	}
	return stringDigits
}

func getFirstAndLast(sliceStringDigits []string) []string {
	firstAndLast := make([]string, len(sliceStringDigits))

	for index, element := range sliceStringDigits {
		firstAndLast[index] = string(element[0]) + string(element[len(element)-1])
	}
	return firstAndLast
}

func convToInt(sliceDigits []string) ([]int, error) {
	intSlice := make([]int, len(sliceDigits))

	for i, value := range sliceDigits {
		int, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		intSlice[i] = int
	}
	return intSlice, nil
}

func sumInts(intSlice []int) int {
	sum := 0

	for _, value := range intSlice {
		sum += value
	}
	fmt.Println(sum)
	return sum
}

func main() {
	fmt.Println("Main!")
	words := fileReader("sample.txt")
	wordToDigitsSlice := wordToDigit(words)
	extractedDigits := extractDigits(wordToDigitsSlice)
	firstAndLast := getFirstAndLast(extractedDigits)
	convertedInts, _ := convToInt(firstAndLast)
	sumInts(convertedInts)
}
