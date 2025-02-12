package semaphore

type Semaphore struct {
    ch chan struct{}
}

// NewSemaphore 初期化
func NewSemaphore(n int) *Semaphore {
    return &Semaphore{ch: make(chan struct{}, n)}
}

// Acquire (ロック取得)
func (s *Semaphore) Acquire() {
    s.ch <- struct{}{}
}

// Release (ロック解放)
func (s *Semaphore) Release() {
    <-s.ch
}
