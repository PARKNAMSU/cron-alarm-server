package common_tool

import (
	"reflect"
	"sync"
)

func ParallelExec(funcs ...func() error) []error {
	var wg *sync.WaitGroup
	var mx *sync.Mutex

	wg.Add(len(funcs))

	var errors []error

	for _, exec := range funcs {
		exec := exec
		go func() {
			defer wg.Done()
			if err := exec(); err != nil {
				mx.Lock()
				errors = append(errors, err)
				mx.Unlock()
			}
		}()
	}
	wg.Wait()
	return errors
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}
