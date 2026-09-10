/*
 * @Author       : lvyitao
 * @Date         : 2024-06-07 15:43:28
 * @LastEditTime: 2026-01-15 09:23:01
 */
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"src/utils"
)

var (
	basePath        = utils.GetExcutePath()
	pid             = -1
	pidFile  string = filepath.Join(basePath, "pid")
)

func initDaemon() {
	ok, _ := utils.PathExists(pidFile)
	if ok {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			log.Println("failed to read pid file", err)
			pid = -1
			return
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Println("failed to parse pid data", err)
			pid = -1
			return
		}
		if utils.IsProcessAlive(id) {
			pid = id
		} else {
			pid = -1
		}
	}
}
