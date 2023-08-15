package utimes

import (
	"sync/atomic"
	"time"
)

var now atomic.Value

func init() {
	refreshOnce()
	go refreshTask()
}

func refreshTask() {
	t := time.NewTicker(time.Millisecond)
	defer t.Stop()
	for {
		<-t.C
		refreshOnce()
	}
}

func refreshOnce() {
	now.Store(time.Now())
}

func CurrentTimeInSec() int64 {
	n := now.Load().(*time.Time)
	return n.UnixNano() / int64(time.Second)
}

func CurrentTimeInMs() int64 {
	n := now.Load().(*time.Time)
	return n.UnixNano() / int64(time.Millisecond)
}

func CurrentTimeInUs() int64 {
	n := now.Load().(*time.Time)
	return n.UnixNano() / int64(time.Microsecond)
}

func CurrentTimeInNs() int64 {
	n := now.Load().(*time.Time)
	return n.UnixNano()
}
