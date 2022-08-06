package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"simple-shell/utils"
	"strings"
)

func main() {
	var reader bufio.Reader = *bufio.NewReader(os.Stdin)
	for {
		out, _ := exec.Command("pwd").Output()
		fmt.Printf("\n %s :) ", strings.TrimSuffix(string(out), "\n"))
		input, _ := reader.ReadString('\n')
		commands_in_sequence := strings.Split(input, ";")
		errs := utils.ParseCommandsSequential(commands_in_sequence...)
		for _, err := range errs {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
