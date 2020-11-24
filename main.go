package main

import (
	"flag"

	"github.com/speier/hrun/pkg/req"
	"github.com/speier/hrun/pkg/utils"
	"github.com/speier/hrun/pkg/vm"
)

var (
	envargs utils.ArrayFlag
)

func init() {
	flag.Var(&envargs, "e", "env var or file (repeatable)")
	flag.Parse()
}

func main() {
	// init vm
	vm := vm.NewInterpreter(req.Methods)
	vm.Set("env", utils.VMEnv(envargs))

	// run script
	_, err := vm.RunFile(flag.Arg(0))
	if err != nil {
		println(err.Error())
	}
}
