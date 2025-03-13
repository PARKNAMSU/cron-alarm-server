package common_tool

import (
	"crypto/rand"
	"math/big"
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

func IsSlice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

// RandomString 함수는 주어진 길이의 랜덤 문자열을 생성합니다.
func RandomString(length int) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}

func Compact[T *any](list []T) []T {
	returnData := make([]T, 0)
	for _, data := range list {
		if data != nil {
			returnData = append(returnData, data)
		}
	}
	return returnData
}
