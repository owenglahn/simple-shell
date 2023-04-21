package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ParseInput(input string) string {
	return CommandSubstitution(input)
}

func CommandSubstitution(input string) string {
	input = strings.TrimSuffix(input, "\n")
	num_ticks := strings.Count(input, "`")
	if num_ticks%2 == 1 {
		// didn't close tick
		fmt.Printf("Didn't close tick")
	}
	backtick_split := strings.Split(input, "`")
	execution_split := make([]string, 0)
	for index, substr := range backtick_split {
		if index%2 == 1 {
			out := BasicExecute(substr)
			execution_split = append(execution_split, out)
		} else {
			execution_split = append(execution_split, substr)
		}
	}
	return strings.Join(execution_split, " ")
}

func BasicExecute(input string) string {
	args := strings.Split(input, " ")
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return string(output)
}

func TrimWhiteSpaceFromAll(input []string) []string {
	for i := range input {
		input[i] = strings.Trim(input[i], " ")
	}
	return input
}
