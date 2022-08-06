package utils

import (
	"os"
	"os/exec"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func ParseInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ParseCommandsParallel(input ...string) []error {
	errors := make([]error, 0)
	input = TrimWhiteSpaceFromAll(input)
	wg.Add(len(input))
	for _, command := range input {
		go func(cmd string) {
			err := ParseInput(cmd)
			errors = append(errors, err)
			wg.Done()
		}(command)
	}
	wg.Wait()
	return errors
}

func ParseCommandsSequential(input ...string) []error {
	errors := make([]error, 0)
	input = TrimWhiteSpaceFromAll(input)
	for _, sequential_input := range input {
		commands_in_parallel := strings.Split(sequential_input, "&")
		errors = append(errors, ParseCommandsParallel(commands_in_parallel...)...)
	}
	return errors
}

func TrimWhiteSpaceFromAll(input []string) []string {
	for i := range input {
		input[i] = strings.Trim(input[i], " ")
	}
	return input
}
