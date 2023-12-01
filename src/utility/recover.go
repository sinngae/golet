package utility

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"
)

func RecoverHandler(ctx context.Context) error {
	if rcv := recover(); rcv != nil {
		tNow := time.Now().Format("2006/01/02 - 15:04:05")
		rcvStr := fmt.Sprintf("%+v", rcv)
		stack := string(debug.Stack())
		err := fmt.Errorf("panic: time=%+v, panic=%+v, stack=%s", tNow, rcvStr, stack)
		return err
	}
	return nil
}
