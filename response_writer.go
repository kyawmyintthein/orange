package orange

import (
	"net/http"
)

type ResponseWriter interface {
	Status() int
	Size() int
	http.ResponseWriter
	http.Flusher
	Before(func(ResponseWriter))
}

type Response struct {
	http.ResponseWriter
	status      int
	size        int
	beforeFuncs []func(ResponseWriter)
}

func (res *Response) Status() int {
	return c.status
}

func (res *Response) Size() int {
	return res.size
}

func (res *Response) Written() bool {
	return res.size != noWritten
}

func (res *Response) WriteHeaderNow() {
	if !res.Written() {
		res.size = 0
		res.callBefore()
		res.ResponseWriter.WriteHeader(c.status)
	}
}

func (res *Response) Before(before func(ResponseWriter)) {
	res.beforeFuncs = append(res.beforeFuncs, before)
}
