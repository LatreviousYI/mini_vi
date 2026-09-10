//go:build linux

package utils

import "syscall"

func IsProcessAlive(pid int) bool {
	// syscall.Kill with signal 0 does not actually send a signal,
	// but it checks for the existence of the process.
	err := syscall.Kill(pid, 0)
	return err == nil
}
