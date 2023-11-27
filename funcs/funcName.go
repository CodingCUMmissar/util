package funcs

import (
	"reflect"
	"runtime"
	"strings"
)

func Name(function interface{}) string {
	switch reflect.ValueOf(function).Kind() {
	case reflect.Func:
		return getFuncName(function)
	default:
		panic("funcname: argument must be a function")
	}
}

func getFuncName(function interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
