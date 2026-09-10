package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func TimeCleanlog() {
	defer ErrCatch()
	CopyLog()
	CleanLog()
}

// cat /dev/null > aaa.log
func CleanLog() {
	basePath := GetExcutePath()
	logfile := filepath.Join(basePath, "server.log")
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cat /dev/null > %s", logfile))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, errStr := stdout.String(), stderr.String()
	if errStr != "" {
		log.Println(errStr)
		return
	}
}

func CopyLog() {
	basePath := GetExcutePath()
	logfile := filepath.Join(basePath, "server.log")
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cp %s %s", logfile, logfile+".bak"))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, errStr := stdout.String(), stderr.String()
	if errStr != "" {
		log.Println(errStr)
		return
	}
}
