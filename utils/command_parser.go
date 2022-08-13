package utils

import (
	"os"
	"os/exec"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func ParseInput(input string) error {
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func CommandSubsitution(input string) error {
	input = strings.TrimSuffix(input, "\n")
	backtick_split := strings.Split(input, "`")
	execution_split := make([]string, 0)
	for index, substr := range backtick_split {
		if index%2 == 1 {
			command_sub_errs := ParseCommandsSequential(substr)
			command_sub_string := ""
			for _, err := range command_sub_errs {
				command_sub_string += err.Error()
			}
			execution_split = append(execution_split, command_sub_string)
		} else {
			execution_split = append(execution_split, substr)
		}
	}
	return ParseInput(strings.Join(execution_split, " "))
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

func ParseCommandsSequential(input string) []error {
	errors := make([]error, 0)
	input = strings.TrimSuffix(input, "\n")
	split_input := TrimWhiteSpaceFromAll(strings.Split(input, ";"))
	for _, sequential_input := range split_input {
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
