package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ArrayFlag []string

func (i *ArrayFlag) String() string {
	return strings.Join(*i, ",")
}

func (i *ArrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func VMEnv(envargs []string) map[string]interface{} {
	env := map[string]interface{}{}

	// 1) load default `.env` file to os.env
	godotenv.Load()

	// 2) try to load files if arg is a file to os.env
	for _, e := range envargs {
		godotenv.Load(e)
	}

	// 3) append rest of the args if valid key=value to os env
	envargs = append(os.Environ(), envargs...)
	for _, e := range envargs {
		k, v := stringkv(e, "=")
		env[k] = v
	}

	return env
}

func stringkv(s string, sep string) (k, v string) {
	a := strings.Split(s, sep)
	if len(a) == 2 {
		k, v = a[0], a[1]
	}
	return
}
