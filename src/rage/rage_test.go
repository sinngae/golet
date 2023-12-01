package rage

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewRage(t *testing.T) {
	rage := NewRage(10)

	job := func(val int) {
		fmt.Printf("do job[%d] ... ...\n", val)
		rnd := rand.Uint32()%3 + 1
		time.Sleep(time.Duration(rnd) * time.Second)
		if rnd == 2 {
			panic("this is a panic for test")
		}
	}

	for idx := 0; idx < 100; idx++ {
		idxVal := idx
		rage.AppendJob(&Job{
			Do: func() {
				job(idxVal)
			},
			PanicHandler: func(pnc interface{}) {
				fmt.Println(pnc)
			},
		})
	}

	rage.Wait()
}
