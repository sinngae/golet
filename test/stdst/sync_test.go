package stdst

import "sync"

func ExampleSync() {
	_ = sync.Mutex{}
	sp := &sync.Pool{}
	a := sp.New()
	sp.Put(a)
	sp.Get()
	//output:
}
