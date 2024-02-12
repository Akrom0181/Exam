package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	var words []string
	for scanner.Scan() {
		line := scanner.Text()

		for _, word := range strings.Fields(line) {
			if num, err := strconv.Atoi(word); err == nil {
				numbers = append(numbers, num)
			} else {
				words = append(words, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scanning:", err)
		return
	}

	numbersFile, err := os.Create("numbersOnly.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer numbersFile.Close()

	var s int
	OnlyNumber := bufio.NewWriter(numbersFile)
	for _, num := range numbers {
		s += num
		fmt.Fprintln(OnlyNumber, num)
	}
	OnlyNumber.Flush()
	fmt.Fprint(numbersFile, s)

	wordsFile, err := os.Create("OnlyWordOutput.txt")
	if err != nil {
		fmt.Println("Error while writing to file: ", err)
		return
	}
	defer wordsFile.Close()

	wordOnly := bufio.NewWriter(wordsFile)
	for _, word := range words {
		if strings.Contains(word, "a") || strings.Contains(word, "o") || strings.Contains(word, "A") || strings.Contains(word, "O") {
			fmt.Fprintln(wordOnly, word)
		}
	}
	wordOnly.Flush()

	fmt.Println("OnlyNumbers: numbersOnly.txt")
	fmt.Println("OnlyWords: OnlyWordOutput.txt")
}
