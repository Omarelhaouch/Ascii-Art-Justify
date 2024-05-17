package ascii

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Define command-line flags
var (
	ColorPtr = flag.String("color", "white", "color for ASCII art (default: white)")
)

// getColorSettings processes the color settings and returns the necessary values
func GetColorSettings() (string, []string, error) {
	color := *ColorPtr

	colors, err := Colors(color)
	if err != nil {
		return "", nil, err
	}

	charsToColor := "" // Default to coloring the whole string
	if len(flag.Args()) > 1 {
		charsToColor = flag.Arg(1) // If a second argument is provided, use it as the characters to color
	}

	return charsToColor, colors, nil

}

func Colors(color string) ([]string, error) {
	if color == "rainbow" {
		// Define a rainbow of colors
		return []string{"\033[31m", "\033[38;5;214m", "\033[33m", "\033[32m", "\033[34m", "\033[38;5;54m", "\033[35m"}, nil
	}
	if strings.HasPrefix(color, "#") {
		// Convert hexadecimal color code to ANSI escape sequence
		// This assumes a 24-bit color terminal
		r, _ := strconv.ParseInt(color[1:3], 16, 64)
		g, _ := strconv.ParseInt(color[3:5], 16, 64)
		b, _ := strconv.ParseInt(color[5:7], 16, 64)
		return []string{fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)}, nil
	}
	if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		// Parse rgb(r, g, b) format
		rgb := strings.TrimPrefix(color, "rgb(")
		rgb = strings.TrimSuffix(rgb, ")")
		parts := strings.Split(rgb, ",")
		if len(parts) != 3 {
			return nil, errors.New("invalid rgb format")
		}
		r, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		g, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		b, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		return []string{fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)}, nil
	}
	switch color {
	case "black":
		return []string{"\033[30m"}, nil
	case "red":
		return []string{"\033[31m"}, nil
	case "green":
		return []string{"\033[32m"}, nil
	case "yellow":
		return []string{"\033[33m"}, nil
	case "blue":
		return []string{"\033[34m"}, nil
	case "magenta":
		return []string{"\033[35m"}, nil
	case "cyan":
		return []string{"\033[36m"}, nil
	case "white":
		return []string{"\033[37m"}, nil
	case "orange":
		return []string{"\033[38;5;214m"}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unknown color: %s", color))
	}
}
