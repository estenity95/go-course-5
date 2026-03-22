package task08concurrency

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add(delta int) {
	c.mu.Lock()
	c.value += delta
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

// IncrementConcurrently много горутин
func IncrementConcurrently(counter *Counter, workers, increments int) {
	var wg sync.WaitGroup
	wg.Add(workers)

	for worker := 0; worker < workers; worker++ {
		go func() {
			defer wg.Done()
			for i := 0; i < increments; i++ {
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
}
