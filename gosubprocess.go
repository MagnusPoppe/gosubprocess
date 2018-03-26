package gosubprocess

import (
	"os/exec"
	"log"
	"io"
	"os"
)

type Program struct {
	command *exec.Cmd
	stdin   io.WriteCloser
	stdout  io.Reader
}

func (program Program) Send(message string) {
	program.stdin.Write([]byte(message + "\n\r"))
	log.Println("Sent message", message)
}

func (program Program) ReadOut() string {
	out := make([]byte, 1024)
	n, _ := program.stdout.Read(out)
	return string(out[:n])
}

func (program Program) Kill() {
	// Closing communication and killing program.
	program.command.Process.Kill()

	// Setting to nil to make sure no errors are made by bad use.
	program.stdout = nil
	program.stdin = nil
	program.command = nil
}

func (program Program) Finish() {
	// Waiting for the program to finish
	program.command.Wait()

	// Setting to nil to make sure no errors are made by bad use.
	program.stdout = nil
	program.stdin = nil
	program.command = nil
}

func SetupProgram(args ...string) Program {

	cmd := exec.Command(args[0], args[1:]...)
	stdin, err := cmd.StdinPipe()
	check(err)
	stdout, err := cmd.StdoutPipe()
	check(err)
	cmd.Stderr = os.Stderr

	cmd.Start()
	return Program{cmd, stdin, stdout}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}