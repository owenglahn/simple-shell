package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-shell/utils"
)

func main() {
	var reader bufio.Reader = *bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = utils.ParseInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
