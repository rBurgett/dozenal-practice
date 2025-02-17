package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rBurgett/dozenal-practice/internal/dozenal"
	"os"
	"strings"
)

func main() {
	maxB10 := flag.String("sum", "10", "The sum to add up to")
	flag.Parse()
	maxI, err := dozenal.FromBase12(*maxB10)
	if err != nil {
		panic(err)
	} else if maxI < 2 {
		fmt.Println("The sum must be greater than 1")
		return
	}

	prev := 0
	for {
		n := dozenal.GetRandom(1, maxI)
		if n == prev {
			prev = 0
			continue
		}
		prev = n
		nn := maxI - n
		fmt.Println(fmt.Sprintf("%s + ? = %s", dozenal.ToBase12(n), dozenal.ToBase12(maxI)))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputNum, err := dozenal.FromBase12(input)
		if err != nil {
			fmt.Println(fmt.Sprintf("Invalid input: %s", input))
			fmt.Println("")
			continue
		}
		if inputNum == nn {
			fmt.Println("Correct!")
		} else {
			fmt.Println(fmt.Sprintf("Incorrect. The answer is %s\n", dozenal.ToBase12(nn)))
		}
		fmt.Println("")
	}
}
