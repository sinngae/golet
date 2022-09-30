package msg_buffer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewMsgBuffer(t *testing.T) {
	msgBuf := NewMsgBuffer(3, 5, func(list []string) {
		fmt.Printf("list: %v\n", list)
	})

	for i := 0; i < 10; i++ {
		msgBuf.Append(fmt.Sprintf("item:%d", i))
		time.Sleep(1 * time.Second)
	}
	msgBuf.Stop()
}
