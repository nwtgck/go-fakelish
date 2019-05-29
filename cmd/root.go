package cmd

import (
	"fmt"
	"github.com/nwtgck/go-fakelish"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var minLength int
var maxLength int
var nWords int
var enableCapitalize bool

func init() {
	cobra.OnInitialize()
	RootCmd.Flags().IntVar(&minLength,  "min", 6, "min length of fake word")
	RootCmd.Flags().IntVar(&maxLength,  "max",  9, "max length of fake word")
	RootCmd.Flags().IntVarP(&nWords,  "n-words",  "n", 10, "number of fake words")
	RootCmd.Flags().BoolVar(&enableCapitalize,  "capitalize", true, "capitalize the first letter")
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "fakelish",
	Long:  "English-like word generator",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < nWords; i++ {
			// Generate a fake word
			fakeWord := fakelish.GenerateFakeWord(minLength, maxLength)
			if enableCapitalize {
				// Capitalize the first letter
				fakeWord = strings.Title(fakeWord)
			}
			// Print the fake word
			fmt.Println(fakeWord)
		}
	},
}
