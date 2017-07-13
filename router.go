package orange

import "github.com/julienschmidt/httprouter"
import "net/http"

type Router struct {
	app          *App
	handlerFuncs []HandlerFunc
	prefix       string
}

func (r *Router) Use(middlewares ...HandlerFunc) {
	r.handlerFuncs = append(r.handlerFuncs, middlewares...)
}

//GET handle GET method
func (r *Router) GET(path string, handlers ...HandlerFunc) {
	r.Handle("GET", path, handlers)
}

//POST handle POST method
func (r *Router) POST(path string, handlers ...HandlerFunc) {
	r.Handle("POST", path, handlers)
}

//PATCH handle PATCH method
func (r *Router) PATCH(path string, handlers ...HandlerFunc) {
	r.Handle("PATCH", path, handlers)
}

//PUT handle PUT method
func (r *Router) PUT(path string, handlers ...HandlerFunc) {
	r.Handle("PUT", path, handlers)
}

//DELETE handle DELETE method
func (r *Router) DELETE(path string, handlers ...HandlerFunc) {
	r.Handle("DELETE", path, handlers)
}

//HEAD handle HEAD method
func (r *Router) HEAD(path string, handlers ...HandlerFunc) {
	r.Handle("HEAD", path, handlers)
}

//OPTIONS handle OPTIONS method
func (r *Router) OPTIONS(path string, handlers ...HandlerFunc) {
	r.Handle("OPTIONS", path, handlers)
}

//Group group route
func (r *Router) Controller(path string, handlers ...HandlerFunc) *Router {
	handlers = r.mergeHandlers(handlers)
	return &Router{
		handlerFuncs: handlers,
		prefix:       r.path(path),
		app:          r.app,
	}
}

//HandlerFunc convert http.HandlerFunc to ace.HandlerFunc
func (r *Router) HTTPHandlerFunc(h http.HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		h(ctx.response, ctx.request)
	}
}

//Handle handle with specific method
func (r *Router) Handle(method, path string, handlers []HandlerFunc) {
	handlers = r.mergeHandlers(handlers)
	r.app.httprouter.Handle(method, r.path(path), func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := r.app.newContext(rw, req)
		ctx.params = params
		ctx.handlerFuncs = handlers
		ctx.Next()
		r.app.pool.Put(ctx)
	})
}

func (r *Router) path(p string) string {
	if r.prefix == "/" {
		return p
	}
	return concat(r.prefix, p)
}

func (r *Router) mergeHandlers(handlers []HandlerFunc) []HandlerFunc {
	aLen := len(r.handlerFuncs)
	hLen := len(handlers)
	h := make([]HandlerFunc, aLen+hLen)
	copy(h, r.handlerFuncs)
	for i := 0; i < hLen; i++ {
		h[aLen+i] = handlers[i]
	}
	return h
}

// ServceHttp
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	r.app.httprouter.ServeHTTP(res, req)
}
