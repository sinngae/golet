package sync

import "sync"

var (
	once sync.Once
	data string
)

func myFun() string {
	once.Do(func() {
		data = "abc"
	})
	return data
}
