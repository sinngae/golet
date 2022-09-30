package stdst

import "sync"

////go:linkname FastRand runtime.fastrand
//func FastRand() uint32

func ExampleMutex() {
	lock := sync.Mutex{}
	lock.Lock()
	lock.Unlock()

	data := sync.Map{}
	data.Store("a", "b")
	//output:
}
