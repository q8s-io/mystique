package sidecar

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var StdinFD = "/proc/4809/fd/0"
var StdoutFD = "/proc/26507/fd/1"
var StderrFD = "/proc/26507/fd/2"

func SinkToStdinFromQueue() {
	f, openErr := os.OpenFile(StdinFD, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if openErr != nil {
		fmt.Println(openErr)
	}

	for msg := range StdinQueue {
		d := fmt.Sprintf("%s\n", string(msg))
		n, writeErr := f.Write([]byte(d))
		if writeErr != nil {
			fmt.Println(writeErr)
		}
		log.Printf("write to %d byte", n)
		_ = f.Sync()
	}

	_ = f.Close()
}

func stdout() {
	file, _ := os.OpenFile(StdoutFD, os.O_RDWR, os.ModeNamedPipe)
	reader := bufio.NewReader(file)
	for {
		line, _, _ := reader.ReadLine()
		fmt.Println("stdout msg", string(line))
	}
}

func stderr() {
	file, _ := os.OpenFile(StderrFD, os.O_RDWR, os.ModeNamedPipe)
	reader := bufio.NewReader(file)
	for {
		line, _, _ := reader.ReadLine()
		fmt.Println("stderr msg", string(line))
	}
}
