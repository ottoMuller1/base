package concurrency



import (
	"sync"
)

// goroutine wrapper for the synchronization control
func GoSync(mutex *sync.Mutex, handle func()) {

	go func() {

		if mutex != nil {
			mutex.Lock()
		}

		handle()

		if mutex != nil {
			mutex.Unlock()
		}

	}()

}
