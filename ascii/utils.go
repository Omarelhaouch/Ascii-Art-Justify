package ascii

import (
	"fmt"
	"os"
	"strings"
)

const (
	GetCharacterHeight = 8
	GetNumOfAsciiXters = 95
)

// ValidateAndDetermineFilename validates the command line arguments and determines the filename of the ASCII art file.
func ValidateAndDetermineFilename() (bool, string) {
    filename := "standard.txt"
    validFilenames := []string{"standard", "shadow", "thinkertoy", "propre"}
    bannerFound := false

    for i, arg := range os.Args {
        for _, validFilename := range validFilenames {
            if arg == validFilename || arg == validFilename+".txt" {
                if bannerFound {
                    fmt.Println("Error: Unexpected argument after banner name.")
                    os.Exit(1)
                }
                filename = validFilename + ".txt"
                bannerFound = true
            }
        }
        if bannerFound && i != len(os.Args)-1 {
            fmt.Println("Error: Unexpected argument after banner name.")
            os.Exit(1)
        }
    }

    return true, filename
}

// ConvertEscapeSequences replaces escape sequences in the input string with their corresponding characters.
func ConvertEscapeSequences(s string) string {
	replacer := strings.NewReplacer(
		"\\v", "\\n\\n\\n\\n",
		"\n", "\\n",
		"\\t", "    ",
		"\\b", "\b",
		"\\r", "\r",
		"\\f", "\f",
	)
	return replacer.Replace(s)
}

// ContainsOnlyPrintableOrWhitespace checks if the input string contains only printable ASCII characters or whitespace.
// RemoveCharactersBeforeBackspace removes characters before a backspace character in the input string.
func RemoveCharactersBeforeBackspace(s string) string {
	temp := ""
	for _, ch := range s {
		if ch != '\b' {
			temp += string(ch)
		} else if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
	}
	return temp
}

func ContainsOnlyPrintableOrWhitespace(args string) bool {
	for _, ch := range args {
		if !((ch >= 32 && ch <= 126) || (ch >= 8 && ch <= 13)) {
			fmt.Println("Error: Input contains invalid characters")
			return false
		}
	}
	return true
}

// CheckForWhitespaceOrNewline checks if the input string array contains only whitespace or newline characters.
func CheckForWhitespaceOrNewline(ar []string) (bool, string, int) {
	count := 0
	for _, ch := range ar {
		if ch != "" {
			return false, "", count
		}
		count++
	}

	if count == 1 {
		return true, "SPACE", count
	}
	return true, "NEWLINE", count
}

// PrintWordInAsciiArt prints a word in ASCII art.
// PrintWordInAsciiArt prints a word in ASCII art.
func PrintWordInAsciiArt(word string, AsciiSymbol map[rune][]string, colorCodes []string, charsToColor string, color string) {
    resetCode := "\033[0m" // ANSI code to reset color

    colorAll := charsToColor == ""

    for i := 0; i < GetCharacterHeight; i++ {
        lineOutput := ""
        beforeCarriage := ""
        for j, l := range word {
            colorCode := colorCodes[j%len(colorCodes)] // Cycle through colors
            switch l {
            case '\n':
                fmt.Println(lineOutput)
                lineOutput = ""
            case '\r':
                beforeCarriage = lineOutput
                lineOutput = ""
            default:
                if i < len(AsciiSymbol[l]) {
                    if colorAll || strings.ContainsRune(charsToColor, l) {
                        lineOutput += colorCode + AsciiSymbol[l][i] + resetCode
                    } else {
                        lineOutput += AsciiSymbol[l][i]
                    }
                }
            }
        }
        if len(beforeCarriage) > len(lineOutput) {
            lineOutput += beforeCarriage[len(lineOutput):]
        }
        fmt.Println(lineOutput)
    }
}