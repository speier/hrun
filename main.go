package main

import (
	"flag"

	"github.com/joho/godotenv"

	"github.com/speier/hrun/pkg/req"
	"github.com/speier/hrun/pkg/utils"
	"github.com/speier/hrun/pkg/vm"
)

var (
	filename string
	params   utils.ArrayFlag
)

func init() {
	godotenv.Load()
	flag.StringVar(&filename, "f", "", "filename")
	flag.Var(&params, "s", "params")
	flag.Parse()
}

func main() {
	// init vm
	vm := vm.NewInterpreter(req.Methods)
	vm.Set("env", utils.VMEnv(params))

	// run script
	_, err := vm.RunFile(filename)
	if err != nil {
		println(err.Error())
	}
}
