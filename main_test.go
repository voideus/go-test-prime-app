package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_alpha_isPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not prime number because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime number, by defination!"},
		{"negative", -11, false, "Negative numbers are not prime number, by defination!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

	// result, msg := isPrime(0)
	// if result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 0)
	// }

	// if msg != "0 is not prime number, by defination!" {
	// 	t.Error("wrong message returned: ", msg)
	// }

	// result, msg = isPrime(7)
	// if !result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 7)
	// }

	// if msg != "7 is a prime number" {
	// 	t.Error("wrong message returned: ", msg)
	// }
}

func Test_alpha_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipt
	os.Stdout = w

	prompt()

	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipt
	os.Stdout = w

	intro()

	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct got: %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime number, by defination!"},
		{name: "two", input: "2", expected: "2 is a prime number"},
		{name: "three", input: "3", expected: "3 is a prime number"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime number, by defination!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "2.2", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this func, we need channel and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
