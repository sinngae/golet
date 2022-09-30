package stdst

import (
	"fmt"
	"testing"

	"github.com/sinngae/golet/src/rage"
)

func TestMapSafe(t *testing.T) {
	defer func() {
		if rcv := recover(); rcv != nil {
			fmt.Println(rcv)
		}
	}()
	data := make(map[int]int)
	//dataLock := sync.Mutex{}

	rg := rage.NewRage(10)

	for i := 0; i < 100; i++ {
		iBk := i
		rg.AppendJob(&rage.Job{
			Do: func() {
				//dataLock.Lock()
				data[iBk] = iBk
				//dataLock.Unlock()
			},
			PanicHandler: nil,
		})
	}

	rg.Wait()

	fmt.Println("test over")
}

func TestMapDel(t *testing.T) {
	data := map[int]int{
		0: 1,
		1: 2,
		9: 2,
		8: 4,
		2: 3,
		3: 0,
	}
	for k, _ := range data {
		if k%2 == 0 {
			delete(data, k)
		}
	}
	println(data)
}

func ExamplePrintMap() {
	myPrint := func(obj interface{}) {
		fmt.Printf("%p %s %v %+v %#v\n", obj, obj, obj, obj, obj)
	}
	data := map[int]int{1: 3, 4: 5, 2: 8}
	myPrint(data)

	//output:

}
