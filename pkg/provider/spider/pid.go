package spider

import (
	"log"
	"os/exec"
	"strings"
)

func runInLinux(cmd string) (string, error) {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(result)), err
	}
}

func GetPid(serverName string) string {
	a := `ps -x | grep "` + serverName + `" | head -1 | awk '{print $1}'`
	pid, err := runInLinux(a)
	if err != nil {
		log.Println(err)
		return ""
	} else {
		return pid
	}
}
