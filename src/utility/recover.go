package utility

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"
)

func RecoverHandler(ctx context.Context) error {
	if rcv := recover(); rcv != nil {
		tNow := time.Now().Format(time.DateTime)
		rcvStr := fmt.Sprintf("%+v", rcv)
		stack := string(debug.Stack())
		err := fmt.Errorf("panic: time=%+v, panic=%+v, stack=%s", tNow, rcvStr, stack)
		return err
	}
	return nil
}
