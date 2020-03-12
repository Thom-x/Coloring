package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "decoration",
	Short: "Execute a program and add prefix suffix and color to its log output",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCommand()
	},
}

var flags struct{
	program string
	prefix string
	suffix string
	color string
	args string
}

var flagsName = struct{
	program, programShort string
	args, argsShort string
	prefix, prefixShort string
	suffix, suffixShort string
	color, colorShort string
} {
	"program", "e",
	"args", "a",
	"prefix", "p",
	"suffix", "s",
	"color", "c",
}

var print func(s string)

func main() {
	rootCmd.Flags().StringVarP(
		&flags.program,
		flagsName.program,
		flagsName.programShort,
		"", "program to execute")
	rootCmd.Flags().StringVarP(
		&flags.args,
		flagsName.args,
		flagsName.argsShort,
		"", "arguments of the program")
	rootCmd.Flags().StringVarP(
		&flags.prefix,
		flagsName.prefix,
		flagsName.prefixShort,
		"", "prefix")
	rootCmd.Flags().StringVarP(
		&flags.suffix,
		flagsName.suffix,
		flagsName.suffixShort,
		"", "suffix")
	rootCmd.Flags().StringVarP(
		&flags.color,
		flagsName.color,
		flagsName.colorShort,
		"", "color : black, red, green, yellow, blue, magenta, cyan, white")

	rootCmd.MarkFlagRequired("program")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}