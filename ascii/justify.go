package ascii

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unicode"
	"unsafe"
)

var (
	align = flag.String("align", "", "text alignment (left, center, right, justify)")
)

const (
	leftAlign    = "left"
	centerAlign  = "center"
	rightAlign   = "right"
	justifyAlign = "justify"
)

func MapFont() map[rune][]string {
	valid, fileName := ValidateAndDetermineFilename()
	if !valid {
		fmt.Println("Invalid filename.")
		os.Exit(1)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	asciiArr := ParseFile(file)

	var asciiStart rune = 32
	ascii := make(map[rune][]string)
	for i, char := range asciiArr {
		ascii[rune(i+int(asciiStart))] = char
	}

	return ascii
}

func ParseFile(file *os.File) [][]string {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var asciiChar []string
	var asciiArr [][]string

	counter := 0
	for fileScanner.Scan() {
		if counter == GetCharacterHeight {
			asciiChar = append(asciiChar, fileScanner.Text())
			asciiArr = append(asciiArr, asciiChar)
			asciiChar = []string{}
			counter = 0
			continue
		}
		counter++
		asciiChar = append(asciiChar, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}

	return asciiArr
}

func TerminalWidth() int {
	var dimensions [2]uint16

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions))); err != 0 {
		fmt.Printf("error getting terminal size: %v\n", err)
	}

	return int(dimensions[1])
}

func PrintOutput(words []string, ascii map[rune][]string, terminalWidth int, align string) {
	var alignment string
	wordsPerLine := 0

	for index, word := range words {
		wordLength := 0

		var beforeCarriage string
		var lineOutput string
		for _, runes := range word {
			switch runes {
			case ' ':
				if align == justifyAlign {
					wordsPerLine++
				}
				wordLength = wordLength + len(ascii[runes][4])
			case '\r':
				beforeCarriage = lineOutput
				lineOutput = ""
			default:
				if len(ascii[runes]) > 4 {
					wordLength = wordLength + len(ascii[runes][4])
				} else {
					fmt.Println("Error: ascii[runes] does not have enough elements.")
					os.Exit(1)
				}
			}
		}
		if len(beforeCarriage) > len(lineOutput) {
			lineOutput += beforeCarriage[len(lineOutput):]
		}
		// fmt.Println(lineOutput)
		fmt.Print(lineOutput)
		if wordLength > terminalWidth {
			fmt.Println("Words don't fit in terminal.")
			os.Exit(0)
		}

		switch align {
		case centerAlign:
			alignment = strings.Repeat(" ", (terminalWidth-wordLength)/2)
		case rightAlign:
			alignment = strings.Repeat(" ", terminalWidth-wordLength)
		case justifyAlign:
			if wordsPerLine == 0 {
				align = "none"
			} else {
				alignment = strings.Repeat(" ", (terminalWidth-wordLength)/wordsPerLine)
			}
		}

		// set i to 1 to equal 8 not 9
		for i := 1; i <= 8; i++ {
			for j, runes := range word {
				if j == 0 && align != justifyAlign {
					fmt.Print(alignment)
				}
				if align == justifyAlign && runes == ' ' {
					fmt.Print(alignment)
				}
				if len(ascii[runes]) > i {
					fmt.Print(ascii[runes][i])
				}
			}
			if i == 8 && index != len(words)-1 {
				continue
			}
			fmt.Println()
		}
		wordsPerLine = 0
	}
}

func IsASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func IsValidAlignment(align string) bool {
	return align == leftAlign || align == centerAlign || align == rightAlign || align == justifyAlign
}
