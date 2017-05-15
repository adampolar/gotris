package main

import (
	"os"
	"os/exec"
)

type Command int

const (
	TIMEDROP Command = iota
	LEFT
	RIGHT
	CLOCKWISE
	ANTICLOCKWISE
	SOFTDROP
	DROP
)

func listenForCommands(channel chan<- Command, stateChannelIn <-chan bool, replyChannel chan<- bool) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	b := make([]byte, 3)
	for {
		select {
		case <-stateChannelIn:
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			replyChannel <- true
			return
		default:
		}
		os.Stdin.Read(b)
		if b[0] != 0 {
			if b[0] == 32 {
				channel <- DROP
			}
			if b[0] == 27 && b[1] == 91 && b[2] < 69 && b[2] > 64 {
				//arrow key
				channel <- []Command{
					CLOCKWISE,
					SOFTDROP,
					RIGHT,
					LEFT,
				}[b[2]-65]
			}
		}
	}
}
