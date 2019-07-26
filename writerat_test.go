package progress

import (
	"io"
	"testing"

	"github.com/matryer/is"
)

type writeAtBufferWrites struct {
	b []byte
	p int64
}

type writeAtBuffer struct {
	writes []*writeAtBufferWrites
}

func (b *writeAtBuffer) WriteAt(p []byte, pos int64) (n int, err error) {
	b.writes = append(b.writes, &writeAtBufferWrites{b: p, p: pos})
	return len(p), nil
}

func TestNewWriterAt(t *testing.T) {
	is := is.New(t)

	// check WriterAt interfaces
	var (
		_ io.WriterAt = (*WriterAt)(nil)
		_ Counter     = (*WriterAt)(nil)
	)

	w := NewWriterAt(&writeAtBuffer{writes: make([]*writeAtBufferWrites, 3)})

	n, err := w.WriteAt([]byte("1"), 0)
	is.NoErr(err)
	is.Equal(n, 1)            // n
	is.Equal(w.N(), int64(1)) // r.N()

	n, err = w.WriteAt([]byte("1"), 0)
	is.NoErr(err)
	is.Equal(n, 1)            // n
	is.Equal(w.N(), int64(2)) // r.N()

	n, err = w.WriteAt([]byte("123"), 2)
	is.NoErr(err)
	is.Equal(n, 3)            // n
	is.Equal(w.N(), int64(5)) // r.N()

}
