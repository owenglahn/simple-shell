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
		fmt.Printf("%s :) ", strings.TrimSuffix(string(out), "\n"))
		input, _ := reader.ReadString('\n')
		errs := utils.ParseCommandsSequential(input)
		for _, err := range errs {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
