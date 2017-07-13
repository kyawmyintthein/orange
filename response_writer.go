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

func (res *Response) Status() int {
	return res.status
}

func (res *Response) Size() int {
	return res.size
}

func (res *Response) Written() bool {
	return res.size != notWritten
}

// func (res *Response) WriteHeader() {
// 	if !res.Written() {
// 		res.size = 0
// 		res.callBefore()
// 		res.ResponseWriter.WriteHeader(res.status)
// 	}
// }

func (res *Response) Before(before func(ResponseWriter)) {
	res.beforeFuncs = append(res.beforeFuncs, before)
}

func (res *Response) WriteHeader(code int) {
	if code > 0 {
		res.status = code
	}
}

func (res *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := res.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (res *Response) CloseNotify() <-chan bool {
	return res.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (res *Response) Flush() {
	flusher, ok := res.ResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}

func (res *Response) callBefore() {
	for i := len(res.beforeFuncs) - 1; i >= 0; i-- {
		res.beforeFuncs[i](res)
	}
}

func (res *Response) reset(writer http.ResponseWriter) {
	res.ResponseWriter = writer
	res.status = http.StatusOK
	res.beforeFuncs = nil
	res.size = notWritten
}
