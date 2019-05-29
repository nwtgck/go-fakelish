package main

import (
	"fmt"
	"github.com/nwtgck/go-fakelish"
)

func main () {
	for i := 0; i < 10; i++ {
		// Generate a fake word
		fakeWord := fakelish.GenerateFakeWord(6, 9)
		// Print the fake word
		fmt.Println(fakeWord)
	}
}
