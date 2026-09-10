/*
 * @Author       : lvyitao
 * @Date         : 2024-06-07 15:56:38
 * @LastEditTime: 2026-01-15 10:31:15
 */
package cmd

import (
	"log"

	"src/common"
	"src/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop() {
	common.GoCronTab.Shutdown()
	initDaemon()
	if pid == -1 {
		log.Println("Seems not have been started. Try use `application start` to start server.")
		return
	}
	err := utils.KillProcessById(pid)
	if err != nil {
		log.Printf("failed to kill process %d: %v", pid, err)
	} else {
		log.Println("stop success killed process: ", pid)
	}
	pid = -1
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop the application",
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}
