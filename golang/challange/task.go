package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type wordTuple struct {
	A string
	B string
}

func main() {
	inputFile := "./input_small.txt"
	// outputFile := "./output_small.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := 0
	inputLinesTotal := 0
	inputLinesCounter := 0
	inputData := []*wordTuple{}
	testData := []*wordTuple{}
	inputMode := true
	for scanner.Scan() {
		// ingore line 0
		if line > 0 {
			if inputLinesCounter < inputLinesTotal {
				if inputMode {
					inputData = append(inputData, &wordTuple{
						A: scanner.Text(),
						B: scanner.Text(),
					})
				} else {
					testData = append(inputData, &wordTuple{
						A: scanner.Text(),
						B: scanner.Text(),
					})
				}
				inputLinesCounter++
				inputMode = !inputMode
			} else {
				inputLinesCounter = 0
				inputLinesTotal, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
				if err != nil {
					panic(fmt.Sprintf("unable to parse input file, line: %v", line))
				}
			}
		}

		// println(scanner.Text())
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine() string {

}
