package main

import (
	"github.com/magnuspoppe/gosubprocess"
	"fmt"
)

func main() {
	// Creating the process and starting the program:
	python := gosubprocess.SetupProgram("python3", "src/example.py")
	defer python.Kill()

	// Reading output from process
	fmt.Printf("Recieved: %s\n", python.ReadOut())

	// Sending a message to the subprocess:
	python.Send("1")

	// Reading output from process
	fmt.Printf("Recieved: %s\n", python.ReadOut())
}
