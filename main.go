package main

import (
	"ascii-art/ascii"
	"flag"
	"fmt"
	// "syscall"
	"log"
	"os"
	"strings"
)

// CheckArgs checks if the required arguments are provided
func CheckArgs() (manip bool, fn string) {
	flag.Parse()
	args := flag.Args()

	// Validate and determine the filename
	manip, fn = ascii.ValidateAndDetermineFilename()
	if !manip {
		// If validation fails, print error message and exit
		fmt.Println("\x1b[38;5;9m\x1b[40mARGS missing EX: go run . [STRING] [BANNER]\nEX: go run . something standard\x1b[38;5;9m\x1b[40m")
		os.Exit(1)
	}

	// Ensure that at least one positional argument is provided
	if len(args) < 1 {
		// If required arguments are missing, print error message and exit
		fmt.Println("\x1b[38;5;9m\x1b[40mARGS missing EX: go run . [STRING] [BANNER]\nEX: go run . something standard\x1b[38;5;9m\x1b[40m")
		os.Exit(1)
	}

	return true, fn
}

// processInput processes the input banner
func processInput() (input []string) {
	param := ascii.ConvertEscapeSequences(flag.Arg(0))
	param = ascii.RemoveCharactersBeforeBackspace(param)
	if !ascii.ContainsOnlyPrintableOrWhitespace(param) {
		// If input contains invalid characters, print error message and exit
		fmt.Println("Input contains invalid characters from banner")
		os.Exit(1)
	}
	input = strings.Split(param, "\\n")
	return
}

// handleWhitespaceOrNewline handles cases where input contains only whitespace or newline characters
func handleWhitespaceOrNewline(input []string) {
	if valid, ifSpace, ind := ascii.CheckForWhitespaceOrNewline(input); valid && ifSpace == "NEWLINE" {
		// If input contains only newline characters, print newlines and exit
		for i := 0; i < ind-1; i++ {
			fmt.Println()
		}
		os.Exit(0)
	} else if valid, ifSpace, _ := ascii.CheckForWhitespaceOrNewline(input); valid && ifSpace == "SPACE" {
		// If input contains only whitespace characters, exit
		os.Exit(0)
	}
}

func main() {
	var color string
	var output string
	var align string

	var flagCount int

	// Custom flag parsing
	for _, arg := range os.Args {
		// Handle --color flag
		switch {
		case strings.HasPrefix(arg, "--color="):
			color = strings.TrimPrefix(arg, "--color=")
			flagCount++
		case strings.HasPrefix(arg, "--color"):
			fmt.Println("Invalid usage of --color. Correct format is --color=<color>")
			os.Exit(1)
		case strings.HasPrefix(arg, "-color"):
			fmt.Println("Invalid usage of --color. Correct format is --color=<color>")
			os.Exit(1)
		}
	
		// Handle --output flag
		switch {
		case strings.HasPrefix(arg, "--output="):
			output = strings.TrimPrefix(arg, "--output=")
			flagCount++
		case strings.HasPrefix(arg, "--output"):
			fmt.Println("Invalid usage of --output. Correct format is --output=<output>")
			os.Exit(1)
		case strings.HasPrefix(arg, "-output"):
			fmt.Println("Invalid usage of --output. Correct format is --output=<output>")
			os.Exit(1)
		}
	
		// Handle --align flag
		switch {
		case strings.HasPrefix(arg, "--align="):
			align = strings.TrimPrefix(arg, "--align=")
			flagCount++
			if align != "left" && align != "right" && align != "center" && align != "justify" {
				fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\nExample: go run . --align=right  something  standard")
				os.Exit(1)
			}
		case strings.HasPrefix(arg, "--align"):
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			os.Exit(1)
		case strings.HasPrefix(arg, "-align"):
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			os.Exit(1)
		}
	}
	
	if flagCount > 1 {
		fmt.Println("Error: Only one flag can be used at a time.")
		os.Exit(1)
	}

	terminalWidth := ascii.TerminalWidth()
	asciiMap := ascii.MapFont()
	align = strings.ToLower(align)
	flag.Parse()

	manip, fn := CheckArgs()
	if !manip {
		os.Exit(1)
	}
	input := processInput()
	handleWhitespaceOrNewline(input)

	new, AsciiSymbol := ascii.LoadAsciiArtFromFile(fn)
	if !new {
		os.Exit(1)
	}

	// Assuming GetColorSettings() uses the `color` variable
	charsToColor, colors, err := ascii.GetColorSettings()
	if err != nil {
		log.Fatal(err)
	}

	// Print each word in the input banner in ASCII art
	// Print each word in the input banner in ASCII art
	for _, str := range input {
		if str == "" {
			fmt.Println()
		} else {
			if output != "" {
				// Open the output file
				// List of banner files
				bannerFiles := []string{"standard.txt", "shadow.txt", "thinkertoy.txt", "propre.txt","main.go"}

				// Check if output file is a banner file or does not have a .txt extension
				for _, bannerFile := range bannerFiles {
					if output == bannerFile || !strings.HasSuffix(output, ".txt") {
						fmt.Println("Error: Cannot write to a banner file or a non-.txt file.")
						os.Exit(1)
					}
				}
				outputFile, err := os.Create(output)
				if err != nil {
					log.Fatal(err)
				}
				defer outputFile.Close()
				ascii.PrintWordInAsciiOutput(str, AsciiSymbol, outputFile)
			} else if color != "" {
				ascii.PrintWordInAsciiArt(str, AsciiSymbol, colors, charsToColor, color)
				// fmt.Printf("!=")
			} else {
				ascii.PrintOutput(strings.Split(str, "\\n"), asciiMap, terminalWidth, align)
			}
		}
	}
}
