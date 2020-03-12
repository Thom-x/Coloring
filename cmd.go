package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"os"
)

func runCommand() error {
	if isInputFromPipe() {
		print("data is from pipe")
		return decorate(os.Stdin, os.Stdout)
	} else {
		file, e := getFile()
		if e != nil {
			return e
		}
		defer file.Close()
		return decorate(file, os.Stdout)
	}
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode() & os.ModeCharDevice == 0
}

func getFile() (*os.File, error){
	if flags.filepath == "" {
		return nil, errors.New("please input a file")
	}
	if !fileExists(flags.filepath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filepath)
	if e != nil {
		return nil, errors.Wrapf(e,
			"unable to read the file %s", flags.filepath)
	}
	return file, nil
}

func decorate(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
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

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
