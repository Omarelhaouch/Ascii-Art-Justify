package ascii

import (
	"fmt"
	"os"
	"flag"
)
var (
	Output = flag.String("output", "", "output file for ASCII art")
)
func PrintWordInAsciiOutput(word string, AsciiSymbol map[rune][]string, outputFile *os.File) {
	// Iterate over each line in the ASCII art
	for i := 0; i < GetCharacterHeight; i++ {
		// Initialize a string to hold the output for the current line
		lineOutput := ""
		// Initialize a string to hold the output before a carriage return
		beforeCarriage := ""
		// Iterate over each character in the word
		for _, l := range word {
			// Switch on the character
			switch l {
			case '\n':
				// If the character is a newline, print the current line output and reset it
				fmt.Fprintln(outputFile, lineOutput)
				lineOutput = ""
			case '\r':
				// If the character is a carriage return, save the current line output and reset it
				beforeCarriage = lineOutput
				lineOutput = ""
			default:
				// If the character is anything else, add its ASCII art to the current line output
				if i < len(AsciiSymbol[l]) {
					lineOutput += AsciiSymbol[l][i]
				}
			}
		}
		// If the output before the carriage return is longer than the current line output,
		// append the extra characters to the current line output
		if len(beforeCarriage) > len(lineOutput) {
			lineOutput += beforeCarriage[len(lineOutput):]
		}
		// Print the current line output
		fmt.Fprintln(outputFile, lineOutput)
	}
}

func ValidateOutput(output string) (outputFileName string, err error) {
	if output == "" {
		err = fmt.Errorf("usage: go run . [OPTION] [STRING] [BANNER]\nexample: go run . --output=<fileName.txt> something standard")
		return
	}
	outputFileName = output
	return
}

// CheckArgs checks if the required arguments are provided and parses the output flag.
func CheckOutput() (manip bool, fn string, outputFileName string) {
    flag.Parse()
    args := flag.Args()
    manip, fn = ValidateAndDetermineFilename()
    if !manip {
        fmt.Println("\x1b[38;5;9m\x1b[40mARGS missing EX: go run . [STRING] [BANNER]\nEX: go run . something standard\x1b[38;5;9m\x1b[40m")
        os.Exit(1)
    }
    if len(args) < 1 {
        fmt.Println("\x1b[38;5;9m\x1b[40mARGS missing EX: go run . [STRING] [BANNER]\nEX: go run . something standard\x1b[38;5;9m\x1b[40m")
        os.Exit(1)
    }
    outputFileName, err := ValidateOutput(*Output)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return true, fn, outputFileName
}