package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "decoration",
	Short: "Transform the input with prefix and suffix",
	Long: `Transform the input with prefix and suffix.
Transform the input (pipe of file) with prefix and suffix`,
	RunE: func(cmd *cobra.Command, args []string) error {
		print = logNoop
		if flags.verbose {
			print = logOut
		}
		return runCommand()
	},
}

var flags struct{
	filepath string
	prefix string
	suffix string
	color string
	verbose bool
}

var flagsName = struct{
	file, fileShort string
	verbose, verboseShort string
	prefix, prefixShort string
	suffix, suffixShort string
	color, colorShort string
} {
	"file", "f",
	"verbose", "v",
	"prefix", "p",
	"suffix", "s",
	"color", "c",
}

var print func(s string)

func main() {
	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the file")
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
	rootCmd.PersistentFlags().BoolVarP(
		&flags.verbose,
		flagsName.verbose,
		flagsName.verboseShort,
		false, "log verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func logNoop(s string) {}

func logOut(s string) {
	log.Println(s)
}