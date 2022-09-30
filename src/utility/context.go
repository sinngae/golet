package utility

import (
	"fmt"
	"time"
)

func WithTimeout(handler func(), timeout time.Duration) error {
	ch := make(chan struct{}, 1)

	go func() {
		handler()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
	case <-time.After(timeout):
		return fmt.Errorf("timeout")
	}
	return nil
}
