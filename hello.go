package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
	"strconv"
	"strings"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number between 1 and 10: ")
	text, _ := reader.ReadString('\n')
	
	text = strings.TrimSpace(text)
	inputNumber, err := strconv.Atoi(text)
	
	if err != nil {
		fmt.Printf("%s is not a number.\n", text)
	} else {
		randomNumber := getRandomNumber(1, 10)
		fmt.Println("The number was: ", randomNumber)
		
		if inputNumber == randomNumber {
			fmt.Println("You were right.")
		} else {
			fmt.Println("You were wrong.")
		}
	}
}

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
