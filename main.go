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
		input = utils.ParseInput(input)
		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
