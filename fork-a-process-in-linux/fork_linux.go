package main

import "syscall"

func fork() (pid int, err error) {
	r1, _, errno := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if errno != 0 {
		return int(r1), errno
	}

	return int(r1), nil
}
