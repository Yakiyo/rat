package cmd

import (
	"bufio"
	"os"
)

func readStdin() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}