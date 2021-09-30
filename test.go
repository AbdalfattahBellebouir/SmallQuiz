package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("problems.csv")
	check(err)
	defer f.Close() // this needs to be after the err check

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Crawlo Quiz\n\nAnswer the following questions:")

	lines, err := csv.NewReader(f).ReadAll()
	check(err)
	correctN := 0
	wrongN := 0

	for _, line := range lines {
		fmt.Print("what is", line[0], " ? -> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		if text == line[1] {
			correctN++
		} else {
			wrongN++
		}
	}
	fmt.Print(correctN, wrongN)
}
