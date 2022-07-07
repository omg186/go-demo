package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	lock := sync.Mutex{}
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer lock.Unlock()
			lock.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter=%d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var lock sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer lock.Unlock()
			lock.Lock()
			counter++
			// lock.Unlock()
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter=%d", counter)
}

func TestCounterThreadGroup(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer lock.Unlock()
			lock.Lock()
			counter++
			wg.Done()
			// lock.Unlock()
		}()
	}
	wg.Wait()
	// time.Sleep(time.Second * 1)
	t.Logf("counter=%d", counter)
}
