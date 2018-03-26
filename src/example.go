package main

import (
	"github.com/magnuspoppe/subprocess"
	"time"
	"fmt"
)

func main() {
	python := subprocess.SetupProgram("python3", "src/example.py")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("Recieved: %s\n", python.ReadOut())
	python.Send("1")
	fmt.Printf("Recieved: %s\n", python.ReadOut())

	python.Terminate()
}
