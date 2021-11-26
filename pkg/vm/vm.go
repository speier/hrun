package vm

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/dop251/goja"
	"github.com/speier/gowasm/pkg/babel"
)

type VM interface {
	Set(name string, value interface{})
	Run(src string) (interface{}, error)
	RunFile(filename string) (interface{}, error)
}

type gojaInterpreter struct {
	runtime *goja.Runtime
}

func NewInterpreter(globals map[string]interface{}) VM {
	runtime := goja.New()
	runtime.SetFieldNameMapper(&jsonTagNamer{})

	// default globals
	runtime.Set("console", map[string]interface{}{
		"log": fmt.Println,
	})
	runtime.Set("expect", func(a, b interface{}) {
		prefix := "OK"
		if a != b {
			prefix = "ERR"
		}
		fmt.Printf("[%s] expected %v got %v\n", prefix, a, b)
	})

	if globals != nil {
		for name, value := range globals {
			runtime.Set(name, value)
		}
	}

	return &gojaInterpreter{runtime}
}

func (i *gojaInterpreter) Set(name string, value interface{}) {
	i.runtime.Set(name, value)
}

func (i *gojaInterpreter) Run(src string) (interface{}, error) {
	es5, err := babel.Transform(src, map[string]interface{}{
		"presets": []string{"latest"},
	})
	if err != nil {
		return nil, err
	}
	return i.runtime.RunString(es5)
}

func (i *gojaInterpreter) RunFile(filename string) (interface{}, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return i.Run(string(src))
}

type jsonTagNamer struct{}

func (*jsonTagNamer) FieldName(t reflect.Type, field reflect.StructField) string {
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		return jsonTag
	}
	return field.Name
}

func (*jsonTagNamer) MethodName(t reflect.Type, method reflect.Method) string {
	return method.Name
}
