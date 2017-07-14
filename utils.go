package orange

import "bytes"
import "sync"

// string concat
func stringConcat(s ...string) string {
	size := 0
	for i := 0; i < len(s); i++ {
		size += len(s[i])
	}

	buf := make([]byte, 0, size)

	for i := 0; i < len(s); i++ {
		buf = append(buf, []byte(s[i])...)
	}

	return string(buf)
}

// Buffer pool
type BufferPool struct{
	pool sync.Pool
}

// NewBufferPool: create new buffer pool
func newBufferPool(size int) *BufferPool {
	var bp BufferPool
	bp.pool.New = func() interface{} {
		return new(bytes.Buffer)
	}
	return &bp
}

// Get: get buffer from pool
func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

// Get: put back buffer to pool
func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	bp.pool.Put(b)
}