package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/speier/hrun/pkg/req"
	"github.com/speier/hrun/pkg/utils"
	"github.com/speier/hrun/pkg/vm"
)

const (
	usage = `Usage: hrun [OPTIONS] <filename>

Options:
`
)

var (
	envargs utils.ArrayFlag
)

func init() {
	flag.Var(&envargs, "e", "env var or file (repeatable)")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	// just append a new line to make output more readable
	fmt.Println()

	// check if file exists
	filename := flag.Arg(0)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Error: file not found '%s'\n", filename)
		fmt.Println()
		flag.Usage()
		os.Exit(0)
	}

	// init vm
	vm := vm.NewInterpreter(req.Methods)
	vm.Set("env", utils.VMEnv(envargs))

	// run script
	_, err := vm.RunFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
}
