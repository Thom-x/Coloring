package main

import (
	"os/exec"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"bufio"
	"os/signal"
	"syscall"
	"github.com/mgutz/str"
)

func runCommand() error {
	return decorate(os.Stdin, os.Stdout)
}


func decorate(r io.Reader, w io.Writer) error {
	args := str.ToArgv(flags.args)
	cmd := exec.Command(flags.program, args...)
	
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		return e
	}

	stderr, e := cmd.StderrPipe()
	if e != nil {
        	return e
    }
	

	e = cmd.Start()
	if e != nil {
		return e
	}

	go func() {
		// wait for the command to finish
		waitCh := make(chan error, 1)
		go func() {
		    waitCh <- cmd.Wait()
		    close(waitCh)
		}()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)

		// Loop to handle multiple signals
		for {
		    select {
		    case sig := <-sigChan:
			cmd.Process.Signal(sig)
		    case err := <-waitCh:
			// Subprocess exited. Get the return code, if we can
			var waitStatus syscall.WaitStatus
			if exitError, ok := err.(*exec.ExitError); ok {
			    waitStatus = exitError.Sys().(syscall.WaitStatus)
			    os.Exit(waitStatus.ExitStatus())
			}
		    }
		}
	}()
	

	merged := io.MultiReader(stdout, stderr)
	scanner := bufio.NewScanner(merged)

	for scanner.Scan() {
		if(flags.color != "") {
			switch(flags.color) {
				case "black":
				color.Set(color.FgBlack)
				case "red":
				color.Set(color.FgRed)
				case "green":
				color.Set(color.FgGreen)
				case "yellow":
				color.Set(color.FgYellow)
				case "blue":
				color.Set(color.FgBlue)
				case "magenta":
				color.Set(color.FgMagenta)
				case "cyan":
				color.Set(color.FgCyan)
				case "white":
				color.Set(color.FgWhite)
			}
		}
		_, e := fmt.Fprintln(w, flags.prefix + scanner.Text() + flags.suffix)
		if(flags.color != "") {
			color.Unset()
		}
		if e != nil {
			return e
		}		
	}
	return nil
}
