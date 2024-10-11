package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker 			run image <cmd> <params>
// go run main.go   run       <cmd< <params>

func main() {

	if len(os.Args) == 1 {
		fmt.Println(os.Args)
		fmt.Println("no command provided")
		return
	}

	switch os.Args[1] {
	case "run":
		run()
	}

}

func run() {
	fmt.Println("Running", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	// we have to the command what the standard out is, otherwise we couldn't see it in the terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/**
Namespaces in Linux:
	- determine what you can see
		- a certain namespace can only see certain processes
*/
