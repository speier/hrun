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

	// 2) try to load file(s) to os.env, if arg is a file
	for _, e := range envargs {
		godotenv.Load(e)
	}

	// 3) append rest of the args if valid key=value to os env
	envargs = append(os.Environ(), envargs...)
	for _, e := range envargs {
		k, v, ok := stringkv(e, "=")
		if ok {
			env[k] = v
		}
	}

	return env
}

func stringkv(s string, sep string) (k string, v string, ok bool) {
	a := strings.Split(s, sep)
	ok = len(a) == 2
	if ok {
		k, v = a[0], a[1]
	}
	return
}
