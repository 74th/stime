package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Printf("stime <command>...\n")
		os.Exit(1)
	}

	var cmd *exec.Cmd
	if len(os.Args) == 2 {
		cmd = exec.Command(os.Args[1])
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}
	stderrPipe, _ := cmd.StderrPipe()

	allOutputWait := new(sync.WaitGroup)
	allOutputWait.Add(1)
	go func() {
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
	allOutputWait.Wait()
	if err == nil {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
