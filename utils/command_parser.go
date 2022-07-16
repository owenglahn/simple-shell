package utils

import (
	"os"
	"os/exec"
	"strings"
)

func ParseInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
