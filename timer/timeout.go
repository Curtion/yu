package timer

import (
	"fmt"
	"sync"
	"time"
)

type timer struct {
	list    map[string]chan any
	mux     *sync.RWMutex
	timeout time.Duration
}

func NewTimer(time time.Duration) *timer {
	return &timer{
		list:    make(map[string]chan any),
		mux:     new(sync.RWMutex),
		timeout: time,
	}
}

func (t *timer) Wait(key string) (any, error) {
	timer := time.NewTimer(t.timeout)
	t.mux.Lock()
	defer t.mux.Unlock()
	t.list[key] = make(chan any)
	select {
	case <-timer.C:
		delete(t.list, key)
		return nil, fmt.Errorf("%v, timeout", key)
	case v := <-t.list[key]:
		delete(t.list, key)
		return v, nil
	}
}

func (t *timer) Done(key string, value any) bool {
	t.mux.Lock()
	defer t.mux.Unlock()
	_, ok := t.list[key]
	if ok {
		t.list[key] <- value
		delete(t.list, key)
	}
	return ok
}
