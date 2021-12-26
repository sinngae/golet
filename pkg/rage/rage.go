package rage

import (
	"fmt"
	debug_2 "github.com/sinngae/gland/pkg/debug"
	"sync"
)

const debug = false

type Rage struct {
	wg sync.WaitGroup

	workerSize     int
	workerCapacity int

	lock sync.Mutex

	jobCh chan *Job
}

type Job struct {
	Do func()

	PanicHandler func(pnc interface{})
}

func NewRage(workerNum int) *Rage {
	if workerNum <= 1 {
		return nil
	}

	return &Rage{
		workerSize:     0,
		workerCapacity: workerNum,

		jobCh: make(chan *Job, 3), // capacity大小不重要
	}
}

// AppendJob Concurrently safe
func (rage *Rage) AppendJob(job *Job) {
	rage.moreWorker()

	rage.jobCh <- job
	rage.wg.Add(1)
}

func (rage *Rage) moreWorker() {
	rage.lock.Lock()
	defer rage.lock.Unlock()
	if rage.workerSize < rage.workerCapacity {
		workerIdx := rage.workerSize
		go func() {
			for item := range rage.jobCh {
				rage.handle(item, workerIdx)
			}
		}()
		rage.workerSize++
	}
}

func (rage *Rage) handle(job *Job, idx int) {
	defer rage.wg.Done()
	defer func() { // MAKE SURE to restart this goroutine when panic
		if r := recover(); r != nil {
			job.PanicHandler(r)
		}
	}()

	if debug_2.IsDebugging(debug) {
		fmt.Printf("worker[%d] ... ...\n", idx)
	}

	job.Do()
}

func (rage *Rage) Wait() {
	rage.wg.Wait()
	close(rage.jobCh)
}
