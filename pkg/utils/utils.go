package utils

import (
	"os"
	"strings"
)

type ArrayFlag []string

func (i *ArrayFlag) String() string {
	return strings.Join(*i, ",")
}

func (i *ArrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func VMEnv(params []string) map[string]interface{} {
	env := map[string]interface{}{}

	for _, e := range os.Environ() {
		k, v := stringkv(e, "=")
		env[k] = v
	}

	for _, p := range params {
		k, v := stringkv(p, "=")
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
