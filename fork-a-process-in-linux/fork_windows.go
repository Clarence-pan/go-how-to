package main

import "errors"

func fork() (pid int, err error) {
	return -1, errors.New("fork() is not supported on windows")
}
