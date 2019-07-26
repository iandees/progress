package progress

import (
	"io"
	"sync"
)

// WriterAt counts the bytes written through it.
type WriterAt struct {
	w io.WriterAt

	lock sync.RWMutex // protects n and err
	n    int64
	err  error
}

// NewWriterAt gets a WriterAt that counts the number
// of bytes written.
func NewWriterAt(w io.WriterAt) *WriterAt {
	return &WriterAt{
		w: w,
	}
}

func (w *WriterAt) WriteAt(p []byte, o int64) (n int, err error) {
	n, err = w.w.WriteAt(p, o)
	w.lock.Lock()
	w.n += int64(n)
	w.err = err
	w.lock.Unlock()
	return
}

// N gets the number of bytes that have been written
// so far.
func (w *WriterAt) N() int64 {
	var n int64
	w.lock.RLock()
	n = w.n
	w.lock.RUnlock()
	return n
}

// Err gets the last error from the WriterAt.
func (w *WriterAt) Err() error {
	var err error
	w.lock.RLock()
	err = w.err
	w.lock.RUnlock()
	return err
}
