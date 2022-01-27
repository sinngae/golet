package stdst

import (
	"fmt"
	"testing"

	"github.com/sinngae/gland/pkg/rage"
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
