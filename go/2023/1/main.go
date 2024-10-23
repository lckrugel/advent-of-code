package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("SUM: ", sumCalibrationValues(*file))
}

func sumCalibrationValues(file os.File) int {
	sum := 0
	reader := bufio.NewReader(&file)
	for {
		if line, _, err := reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			lineDigits := getLineDigits(line)
			n := digitsBytesToInt(lineDigits)
			sum += n
		}
	}
	return sum
}

func getLineDigits(bytes []byte) []byte {
	allLineDigits := make([]byte, 0, 3)
	for _, c := range bytes {
		if c >= 48 && c <= 57 {
			allLineDigits = append(allLineDigits, c)
		}
	}

	firstAndLastDigits := make([]byte, 0, 2)
	firstAndLastDigits = append(firstAndLastDigits, allLineDigits[0], allLineDigits[len(allLineDigits)-1])
	return firstAndLastDigits
}

func digitsBytesToInt(digits []byte) int {
	intSlice := make([]int, 0, 2)

	// Transforma os bytes em um slice de int
	for _, c := range digits {
		intSlice = append(intSlice, int(c)-48)
	}

	result := 0
	// Transforma o slice de inteiros em um número
	for i, v := range intSlice {
		result += v * (powInt(10, len(intSlice)-i-1))
	}
	return result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
