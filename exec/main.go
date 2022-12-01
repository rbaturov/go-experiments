package main

import (
	"log"
	"syscall"
)

const (
	python3Path = "/usr/bin/python3"
	runBinPath  = "/home/titzhak/go/code/local/go-experiments/exec/pythonprog/run"
)

func main() {
	argv := []string{
		"run",
		"bla1",
		"bla2",
		"bla4",
	}
	err := syscall.Exec(runBinPath, argv, []string{})
	if err != nil {
		log.Fatal(err)
	}
}
