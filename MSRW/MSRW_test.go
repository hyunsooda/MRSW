package MSRW

import (
	"fmt"
	"runtime"
	"time"
	"testing"
)

func ReadGlobalData() int {
	b := a
	fmt.Println("READ")
	return b
}

func WriteGlobalData() {
	a = 2
	fmt.Println("WRITE")
}

var a int = 1

func TestRun(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	locker := NewLocker()

	// TODO : Performance Check
	go func() {
		for {
			locker.ReadLock()
			ReadGlobalData()
			locker.ReadUnlock()
		}
	}()

	go func() {
		for {
			locker.WriteLock()
			WriteGlobalData()
			locker.WriteUnlock()
			time.Sleep(time.Second)
		}
	}()

	for {
	}
}
