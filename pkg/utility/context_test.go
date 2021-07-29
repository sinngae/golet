package utility

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithTimeout(t *testing.T) {
	t.Run("test timeout", func(t *testing.T) {
		err := WithTimeout(func() { time.Sleep(4 * time.Second) }, 3*time.Second)
		assert.Errorf(t, err, "timeout")
	})

	t.Run("test done", func(t *testing.T) {
		err := WithTimeout(func() { time.Sleep(3 * time.Second) }, 4*time.Second)
		assert.NoError(t, err)
	})
}

func TestWithTimeout2(t *testing.T) {
	t.Run("test timeout", func(t *testing.T) {
		err := WithTimeout(func() {
			for idx := 0; idx < 30; idx++ {
				time.Sleep(1 * time.Second)
				fmt.Printf("do something and cost %d second\n", idx)
			}
		}, 5*time.Second)
		fmt.Printf("got err=%v\n", err)
		time.Sleep(30 * time.Second)
		assert.Errorf(t, err, "timeout")
	})
}

func TestWithTimeout3(t *testing.T) {
	t.Run("test timeout", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			err := WithTimeout(func() {
				time.Sleep(1 * time.Second)
				//fmt.Printf("do something and cost 1 second\n")
			}, 1*time.Millisecond)
			assert.Errorf(t, err, "timeout")
		}
		time.Sleep(time.Second * 2)
		t.Log(runtime.NumGoroutine())
	})
}
