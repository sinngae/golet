<<<<<<< HEAD
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
=======
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
>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511
