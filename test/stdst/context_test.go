package stdst

import (
	"context"
	"fmt"
	"time"
)

func TaskToCancel(ctx context.Context) {
	for _, item := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} {
		time.Sleep(1 * time.Second) // mocking do something
		fmt.Printf("do somthing... %d\n", item)
		select {
		case <-ctx.Done():
			fmt.Printf("timeout ...\n")
			return // break to end
		default: // non-blocking
			fmt.Printf("skip ...\n")
		}
	}
}

func TaskMayTimeout(ctx context.Context) {
	ctxT, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	go func() {
		time.AfterFunc(time.Second*3, func() { // 本来 5s后跳出，现在3秒
			cancel()
		})
	}()
	TaskToCancel(ctxT)

}

func CtxDemo() {
	ctxB := context.Background()
	ctx, stop := context.WithCancel(ctxB) // , time.Second*3
	defer stop()
	//i := 1

	TaskMayTimeout(ctx)
}

func ExampleContext() {
	CtxDemo()
	fmt.Printf("progress end...\n")
	//output:

}
