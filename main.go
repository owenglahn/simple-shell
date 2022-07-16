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
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = utils.ParseInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
