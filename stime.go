package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Printf("stime <command>...\n")
		os.Exit(1)
	}
	start := time.Now()

	var err error
	if len(os.Args) == 2 {
		err = exec.Command(os.Args[1]).Run()
	} else {
		err = exec.Command(os.Args[1], os.Args[2:]...).Run()
	}
	if err != nil {
		fmt.Printf("output error\n")
	}

	duration := time.Now().Sub(start)
	fmt.Printf("%.3f", float64(duration)/float64(time.Second))
	if err == nil {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
