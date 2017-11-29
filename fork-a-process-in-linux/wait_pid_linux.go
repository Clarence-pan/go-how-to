package main

import "syscall"

func waitPid(pid int) error {
	syscall.Wait4(pid, nil, 0, nil)
	return nil
}
