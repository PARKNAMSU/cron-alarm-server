package common_tool

import (
	"reflect"
	"sync"
)

func ParallelExec(funcs ...func()) {
	var wg *sync.WaitGroup

	wg.Add(len(funcs))

	for _, exec := range funcs {
		exec := exec
		go func() {
			defer wg.Done()
			exec()
		}()
	}
	wg.Wait()
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}
