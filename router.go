package orange

import "github.com/julienschmidt/httprouter"
import "net/http"

type Router struct {
	app          *App
	handlerFuncs []HandlerFunc
	name         string
	absolutePath string
}

func (ns *Router) Use(middlewares ...HandlerFunc) {
	c.handlerFuncs = append(c.handlerFuncs, middlewares...)
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
func (r *Router) Namespace(path string, handlers ...HandlerFunc) *Router {
	handlers = r.mergeHandlers(handlers)
	return &Router{
		handlerFuncs: handlers,
		name:         r.path(path),
		app:          r.app,
	}
}

// //RouteNotFound call when route does not match
// func (r *Router) RouteNotFound(h HandlerFunc) {
// 	r.app.notfoundFunc = h
// }

// //Panic call when panic was called
// func (r *Router) Panic(h PanicHandler) {
// 	r.app.panicFunc = h
// }

//HandlerFunc convert http.HandlerFunc to ace.HandlerFunc
func (r *Router) HTTPHandlerFunc(h http.HandlerFunc) HandlerFunc {
	return func(c *context) {
		h(c.Writer, c.Request)
	}
}

// //Static server static file
// //path is url path
// //root is root directory
// func (r *Router) Static(path string, root http.Dir, handlers ...HandlerFunc) {
// 	path = r.path(path)
// 	fileServer := http.StripPrefix(path, http.FileServer(root))

// 	handlers = append(handlers, func(c *C) {
// 		fileServer.ServeHTTP(c.Writer, c.Request)
// 	})

// 	r.ace.httprouter.Handle("GET", r.staticPath(path), func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 		c := r.ace.createContext(w, req)
// 		c.handlers = handlers
// 		c.Next()
// 		r.ace.pool.Put(c)
// 	})
// }

//Handle handle with specific method
func (r *Router) Handle(method, path string, handlers []HandlerFunc) {
	handlers = r.mergeHandlers(handlers)
	r.app.httprouter.Handle(method, r.path(path), func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
		c := r.app.newContext(rw, req)
		c.params = params
		c.handlers = handlers
		c.Next()
		r.ace.pool.Put(c)
	})
}

func (r *Router) path(p string) string {
	if r.prefix == "/" {
		return p
	}
	return concat(r.prefix, p)
}

func (r *Router) mergeHandlers(handlers []HandlerFunc) []HandlerFunc {
	aLen := len(r.handlers)
	hLen := len(handlers)
	h := make([]HandlerFunc, aLen+hLen)
	copy(h, r.handlers)
	for i := 0; i < hLen; i++ {
		h[aLen+i] = handlers[i]
	}
	return h
}
