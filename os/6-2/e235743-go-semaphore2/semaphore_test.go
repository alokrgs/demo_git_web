package semaphore

import (
    "sync"
    "testing"
)

func TestSemaphore(t *testing.T) {
    s := NewSemaphore(2)
    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        s.Acquire()
        defer s.Release()
        wg.Done()
    }()

    go func() {
        s.Acquire()
        defer s.Release()
        wg.Done()
    }()

    go func() {
        s.Acquire()
        defer s.Release()
        wg.Done()
    }()

    wg.Wait()
}
