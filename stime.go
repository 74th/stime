package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("stime <command>...\n")
		os.Exit(1)
	}

	var cmd *exec.Cmd
	if len(os.Args) == 2 {
		cmd = exec.Command(os.Args[1])
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}

	allOutputWait := new(sync.WaitGroup)
	allOutputWait.Add(1)
	stderrPipe, _ := cmd.StderrPipe()
	go func() {
		// write stderr
		for {
			_, err := io.Copy(os.Stderr, stderrPipe)
			if err != nil {
				allOutputWait.Done()
				return
			}
		}
	}()

	startTime := time.Now()
	cmd.Start()

	err := cmd.Wait()
	duration := time.Now().Sub(startTime)
	fmt.Printf("%.3f", float64(duration)/float64(time.Second))

	// wait for writing all stderr
	allOutputWait.Wait()

	if err == nil {
		os.Exit(0)
	} else {
		// get exit code
		// https://qiita.com/hnakamur/items/5e6f22bda8334e190f63
		if exitErr, ok := err.(*exec.ExitError); ok {
			if s, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				os.Exit(s.ExitStatus())
			}
		}
		os.Exit(1)
	}
}
