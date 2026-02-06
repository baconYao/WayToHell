package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadMove reads two coordinates (from, to) from stdin, space-separated.
// Valid range for each is [0, maxIndex] (e.g. maxIndex 29 for world size 30).
// Returns (from, to, nil) or (0, 0, error) on invalid input.
func ReadMove(maxIndex int) (from, to int, err error) {
	fmt.Print("Enter from and to (e.g. 1 5): ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if scanner.Err() != nil {
			return 0, 0, scanner.Err()
		}
		return 0, 0, errors.New("no input")
	}
	line := strings.TrimSpace(scanner.Text())
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("need exactly two numbers, got %d", len(parts))
	}
	from, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid from: %w", err)
	}
	to, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid to: %w", err)
	}
	if from < 0 || from > maxIndex {
		return 0, 0, fmt.Errorf("from must be in [0, %d], got %d", maxIndex, from)
	}
	if to < 0 || to > maxIndex {
		return 0, 0, fmt.Errorf("to must be in [0, %d], got %d", maxIndex, to)
	}
	return from, to, nil
}
