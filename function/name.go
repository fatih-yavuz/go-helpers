package function

import "runtime"

import (
	"reflect"
)

func Name(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
