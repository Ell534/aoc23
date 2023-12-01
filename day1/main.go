package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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
		// fmt.Println(scanner.Text())
		scanSlice = append(scanSlice, scanner.Text())
	}
	//	fmt.Println(len(scanSlice))
	return scanSlice
}

func digitExtract(stringSlice []string) []string {

	var digitSlice = make([]string, len(stringSlice))
	for wordIndex, word := range stringSlice {
		for _, char := range word {
			if unicode.IsDigit(char) {
				digitSlice[wordIndex] = digitSlice[wordIndex] + string(char)
			}
		}
	}

	for i, word := range digitSlice {
		if len(word) == 1 {
			digitSlice[i] += digitSlice[i]
		}
	}
	// fmt.Println(digitSlice)
	// fmt.Printf("%T\n", digitSlice[0])
	return digitSlice
}

func twoDigits(stringSlice []string) []string {
	fmt.Println("hello from twoDigits")
	for i, digits := range stringSlice {
		if len(digits) > 2 {
			stringSlice[i] = string(digits[0]) + string(digits[len(digits)-1])
		}
	}
	// fmt.Println(stringSlice)
	return stringSlice
}

func strToInt(stringSlice []string) int {

	var sumOfValues = 0

	for _, digits := range stringSlice {
		intValue, err := strconv.Atoi(digits)
		if err != nil {
			fmt.Printf("Atoi error:\n %d", err)
		}
		sumOfValues += intValue
	}
	fmt.Println("Hopefully this works!")
	fmt.Println(sumOfValues)
	return sumOfValues
}

func main() {
	textSlice := fileReader("sample.txt")
	digitSlice := digitExtract(textSlice)
	twoDigitSlice := twoDigits(digitSlice)
	strToInt(twoDigitSlice)
}
