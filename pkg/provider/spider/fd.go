package spider

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SinkToStdinFromQueue(pid string) {
	stdinFD := fmt.Sprintf("/proc/%s/fd/0", pid)

	f, openErr := os.OpenFile(stdinFD, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if openErr != nil {
		fmt.Println(openErr)
	}

	for msg := range StdinQueue {
		d := fmt.Sprintf("%s\n", string(msg))
		n, writeErr := f.Write([]byte(d))
		if writeErr != nil {
			fmt.Println(writeErr)
		}
		log.Printf("write data %d byte to fd.", n)
		_ = f.Sync()
	}

	_ = f.Close()
}

func SinkToQueueFromStdout(pid string) {
	stdoutFD := fmt.Sprintf("/proc/%s/fd/1", pid)

	file, _ := os.OpenFile(stdoutFD, os.O_RDWR, os.ModeNamedPipe)

	reader := bufio.NewReader(file)

	for {
		line, _, _ := reader.ReadLine()
		go processOutputLine(string(line))
	}
}

func stderr(pid string) {
	stderrFD := fmt.Sprintf("/proc/%s/fd/0", pid)

	file, _ := os.OpenFile(stderrFD, os.O_RDWR, os.ModeNamedPipe)

	reader := bufio.NewReader(file)

	for {
		line, _, _ := reader.ReadLine()
		fmt.Println("stderr msg", string(line))
	}
}

func processOutputLine(line string) {
	StdoutQueue <- line
	log.Printf("read data from fd %s.", string(line))
}
