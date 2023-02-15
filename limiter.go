package limiter

import "sync"

type Limiter struct {
	pool chan struct{}
	wg   sync.WaitGroup
}

func New(n int) *Limiter {
	return &Limiter{pool: make(chan struct{}, n), wg: sync.WaitGroup{}}
}

func (l *Limiter) Run(task func()) {
	l.wg.Add(1)
	l.pool <- struct{}{}
	go func() {
		task()
		<-l.pool
		l.wg.Done()
	}()
}

func (l *Limiter) Wait() {
	l.wg.Wait()
}
