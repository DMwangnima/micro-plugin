package sync

import "sync"

type (
	SharedCalls interface {
		Do(key string, fn func() (interface{}, error)) (interface{}, error)
		DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
	}

	call struct {
		wg  sync.WaitGroup
		val interface{}
		err error
	}

	calls struct {
		mu sync.Mutex
		m  map[string]*call
	}
)

func (cs *calls) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	c, done := cs.createCall(key)
	if done {
		return c.val, c.err
	}

	cs.makeCall(c, key, fn)
	return c.val, c.err
}

func (cs *calls) DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error) {
	c, done := cs.createCall(key)
	if done {
		return c.val, false, c.err
	}

	cs.makeCall(c, key, fn)
	return c.val, true, c.err
}

func (cs *calls) createCall(key string) (c *call, done bool) {
	cs.mu.Lock()
	if c, ok := cs.m[key]; ok {
		cs.mu.Unlock()
		c.wg.Wait()
		return c, true
	}

	c = new(call)
	c.wg.Add(1)
	cs.m[key] = c
	cs.mu.Unlock()
	return c, false
}

func (cs *calls) makeCall(c *call, key string, fn func() (interface{}, error)) {
	c.val, c.err = fn()
	cs.mu.Lock()
	delete(cs.m, key)
	cs.mu.Unlock()
	c.wg.Done()
}

func NewSharedCalls() SharedCalls {
	return &calls{
		m:  make(map[string]*call),
	}
}
