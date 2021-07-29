package msg_buffer

import (
	"fmt"
	"sync"
	"time"

	"github.com/sinngae/gland/pkg/internal/debug_"
)

const debug = false

type MsgBuffer struct {
	data  []string
	lock  sync.Mutex
	limit int

	handler func([]string)

	timer    *time.Timer
	interval time.Duration

	nextSendTime time.Time
	keepBufSecs  time.Duration

	control chan int
}

const ( // control enum
	ctlStop = iota + 1
	ctlOver
)

// NewMsgBuffer MUST: keepBufSecs > 1
func NewMsgBuffer(size int, keepBufSecs int, handler func([]string)) *MsgBuffer {
	if size <= 0 || keepBufSecs <= 1 || handler == nil {
		return nil
	}

	keepBufTime := time.Duration(keepBufSecs) * time.Second
	interval := 1 * time.Second
	nextSendTime := time.Now().Add(keepBufTime)
	obj := &MsgBuffer{
		data:  []string{},
		limit: size,

		handler: handler,

		timer:    time.NewTimer(interval),
		interval: interval,

		nextSendTime: nextSendTime,
		keepBufSecs:  keepBufTime,

		control: make(chan int),
	}

	go func() {
		defer obj.timer.Stop()
		func() {
			for {
				select {
				case <-obj.timer.C:
					obj.tick(true)
					obj.timer.Reset(interval)
				case ctl := <-obj.control:
					if ctl == ctlStop {
						if debug_.IsDebugging(debug) {
							fmt.Printf("msg_buffer: stop ...\n")
						}
						return
					}
				}
			}
		}()
		obj.tick(false) // tick at last
		obj.control <- ctlOver
	}()

	return obj
}

func (obj *MsgBuffer) Append(item string) {
	obj.lock.Lock()
	defer obj.lock.Unlock()
	obj.data = append(obj.data, item)
	//obj.timer.Reset(1 * time.Second) // 1s后
}

func (obj *MsgBuffer) handle(now time.Time, running bool) {
	obj.lock.Lock()
	defer obj.lock.Unlock()

	if len(obj.data) <= 0 {
		return
	}

	if running {
		if len(obj.data) <= obj.limit && now.Before(obj.nextSendTime) {
			if debug_.IsDebugging(debug) {
				fmt.Printf("@debug msg_buffer: wait buffer full or timer tick ...\n")
			}
			return
		}
	}

	obj.handler(obj.data)

	obj.data = obj.data[:0] // reset nil
	obj.nextSendTime = now.Add(obj.keepBufSecs)
}

func (obj *MsgBuffer) tick(running bool) {
	now := time.Now()
	nowStr := now.Format(time.RFC3339)
	if running {
		if debug_.IsDebugging(debug) {
			fmt.Printf("msg_buffer: timer tick @%s ...\n", nowStr)
		}
	}

	obj.handle(now, running)
}

func (obj *MsgBuffer) Stop() {
	obj.control <- ctlStop
	for {
		ctl := <-obj.control // wait for control done
		if ctl == ctlOver {
			break
		}
	}
}
