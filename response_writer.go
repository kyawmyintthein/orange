package orange

import (
	"net/http"
	"net"
	"bufio"
	"errors"
)

const notWritten = -1
type ResponseWriter interface {
	Status() int
	Size() int
	Written() bool
	http.ResponseWriter
	http.Flusher
	Before(func(ResponseWriter))
}

type Response struct {
	http.ResponseWriter
	status      int
	size        int
	beforeFuncs []func(ResponseWriter)
	app         *App
}


// Status: return response status
func (res *Response) Status() int {
	return res.status
}

// Size: return response size
func (res *Response) Size() int {
	return res.size
}

// Written: return response is written
func (res *Response) Written() bool {
	return res.size != notWritten
}

// Before: Before allows for a function to be called before the ResponseWriter has been written.
func (res *Response) Before(before func(ResponseWriter)) {
	res.beforeFuncs = append(res.beforeFuncs, before)
}

// WriteHeader: implement http.Handler function write header
func (res *Response) WriteHeader(code int) {
	if res.Written() {
		colorLog("[WARN] Headers were already written!")
	}
	res.size = 0
	res.status = code
	res.callBefore()
	res.ResponseWriter.WriteHeader(res.status)
}

// Hijack:
func (res *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := res.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

// func (res *Response) CloseNotify() <-chan bool {
// 	return res.ResponseWriter.(http.CloseNotifier).CloseNotify()
// }

func (res *Response) Flush() {
	flusher, ok := res.ResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}

// callBefore: call middlewares before written.
func (res *Response) callBefore() {
	for i := len(res.beforeFuncs) - 1; i >= 0; i-- {
		res.beforeFuncs[i](res)
	}
}

// reset: reset response writer
func (res *Response) reset(writer http.ResponseWriter) {
	res.ResponseWriter = writer
	res.status = http.StatusOK
	res.beforeFuncs = nil
	res.size = notWritten
}
