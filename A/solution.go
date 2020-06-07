package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input-201.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	var result int = 0
	var line string
	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		// fmt.Printf(" > Read %d characters\n", len(line))

		i, err := strconv.Atoi(strings.Replace(line, "\n", "", 1))
		if err != nil {
			// handle error
			// fmt.Println(err)
			os.Exit(2)
		}
		// fmt.Println(line, i)

		result = result ^ i
		// fmt.Println(result)
	}

	fmt.Println(result)
}
