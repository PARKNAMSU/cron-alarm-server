package common_tool

import "sync"

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
