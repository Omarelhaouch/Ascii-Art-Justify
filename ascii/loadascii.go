package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func LoadAsciiArtFromFile(filename string) (bool, map[rune][]string) {
	file, err := os.Open(filename)
	if err != nil {
		return false, nil
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// Initialize the map to hold the ASCII art
	AsciiSymbol := make(map[rune][]string)
	// Initialize the current letter and count
	var letter rune = ' '
	var count int
	// Scan through the file line by line
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()
		// If the count is not equal to the character height, append the line to the current letter's ASCII art
		if count != GetCharacterHeight {
			AsciiSymbol[letter] = append(AsciiSymbol[letter], line)
			count++
		} else {
			// If the count is equal to the character height, move on to the next letter and reset the count
			letter++
			count = 0
		}
		// If there's an error scanning the file, print an error message and exit
		if err := scanner.Err(); err != nil {
			fmt.Printf("\x1b[38;5;9m[Internal Error] Error reading file: %v\x1b[38;5;9m", err)
			os.Exit(1)
		}
	}

	// If the number of characters in the ASCII art does not match the expected number, print an error message and exit
	if len(AsciiSymbol) != GetNumOfAsciiXters {
		fmt.Printf("\x1b[38;5;9m[Error] Expected %d but got %d characters in the alphabet. Ensure you have the correct ascii art file\n\x1b[38;5;9m", GetNumOfAsciiXters, len(AsciiSymbol))
		os.Exit(1)
	}
	// If everything went well, return true and the ASCII art
	return true, AsciiSymbol
}
