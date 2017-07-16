package orange

import "bytes"

// Buffer pool
type BufferPool struct{
	buffer chan *bytes.Buffer
}

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

// NewBufferPool: create new buffer pool
func newBufferPool(size int) *BufferPool {
	return &BufferPool{
		buffer: make(chan *bytes.Buffer, size),
	}
}

// Get: get buffer from pool
func (bp *BufferPool) Get() (b *bytes.Buffer) {
	select {
	case b = <-bp.buffer:
	default:
		b = bytes.NewBuffer([]byte{})
	}
	return
}

// Get: put back buffer to pool
func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	select {
	case bp.c <- b:
	default: 
	}
}