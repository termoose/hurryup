package tools

import "sync"

type Chan[T any] struct {
	Data chan T
	lock sync.Mutex
}

func NewChan[T any]() Chan[T] {
	return Chan[T]{
		Data: make(chan T),
	}
}

func (c *Chan[T]) Send(elem T) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Data <- elem
}

func (c *Chan[T]) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()

	close(c.Data)
}
