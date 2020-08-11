package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("stdout", scanner.Text())
		_, _ = fmt.Fprintln(os.Stderr, "stderr", scanner.Text())
	}
}
