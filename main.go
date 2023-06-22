package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	n := 2

	_, msg := isPrime(n)
	fmt.Println(msg)
	// // print a welcome message
	// intro()

	// // create a channel to indicate when the user wants to quit
	// doneChan := make(chan bool)

	// // start a goroutine to read user input and run program
	// go readUserInput(os.Stdin, doneChan)

	// // block until the doneChan gets a value
	// <-doneChan

	// // close the channel
	// close(doneChan)

	// // say goodbye
	// fmt.Println("Goodbye !!!")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wnats to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert user input to int
	numToCheck, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("--------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by defination
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime number, by defination!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime number, by defination!"
	}

	// use the modulus operator to see if a number is a prime
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
