package orange

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"sync"
)

// Error
var (
	notFoundError       = NewHttpError(http.StatusNotFound)
	internalServerError = NewHttpError(http.InternalServerError)
)

// bufffer pool
var bufPool = pool.NewBufferPool(100)

type HandlerFunc func(ctx *Context)
type config *viper.Viper
type App struct {
	router *Router
	router *httprouter.Router
	pool   sync.Pool
}

// New: init new app object
func New() *App {
	var app *App
	app = new(App)
	app.defaultPool()
	app.newRouter()
	return app
}

// defaultPool: load default pool
func (app *App) defaultPool() {
	app.pool{
		New: func() interface{} {
			return app.newContext()
		},
	}
}

// newContext: init new context for each request and response
func (app *App) newContext(rw http.ResponseWriter, req *http.Request) *Context {
	var ctx *context
	ctx = new(context)
	ctx := app.pool.Get().(*context)
	ctx.Request = req
	ctx.response = &Response{app: app, Writer: rw}
	ctx.index = -1
	ctx.data = nil
	ctx.app = app
	return ctx
}

// newRouter: init new router
func (app *App) newRouter() {
	var (
		httprouter *httprouter.Router
	)
	a.Router = &Router{
		handlers: nil,
		prefix:   "/",
		app:      app,
	}
	httprouter = httprouter.New()
	app.httprouter = httprouter
	app.handleNotFound()
	app.handlePanic()
}

// handleNotFound:  hanlder function for not found
func (app *App) handleNotFound() {
	app.router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx *context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s", ContentTypeJSON, UTF8))
		ctx.response.WriteHeader(http.StatusNotFound)
		ctx.Next()

		buf := bufPool.Get()
		defer bufPool.Put(buf)

		if err := json.NewEncoder(buf).Encode(notFoundError); err != nil {
			return
		}
		c.response.Write(buf.Bytes())
		a.pool.Put(c)
	})
}

// handlePanic: handler function for panic
func (app *App) handlePanic() {
	app.router.PanicHandler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx *context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s", ContentTypeJSON, UTF8))
		ctx.response.WriteHeader(http.InternalServerError)
		ctx.Next()

		buf := bufPool.Get()
		defer bufPool.Put(buf)

		if err := json.NewEncoder(buf).Encode(internalServerError); err != nil {
			return
		}
		c.response.Write(buf.Bytes())
		a.pool.Put(ctx)
	})
}

// ServceHttp
func (app *App) ServceHttp(res http.ResponseWriter, req *http.Request) {
	app.router.ServeHttp(res, req)
}

// Start: start http server
func (app *App) Start(addr stirng) {
	ColorLog("[INFO] server start at: %s\n", addr)
	if err := http.ListenAndServe(addr, app.router); err != nil {
		panic(err)
	}
}

// Start lts (https) server
func (app *App) StartTLS(addr string, cert string, key string) {
	if err := http.ListenAndServeTLS(addr, cert, key, app.router); err != nil {
		panic(err)
	}
}

func (app *App) Namespace(path string, handlers ...HandleFunc, middlewares ...HandleFunc) *Router {
	handlers = app.router.mergeHandlers(handlers)
	router := Router{
		handlerFuncs: handlers,
		name:         r.path(path),
		app:          r.app,
	}

	router.Use(middlewares)
	return &router
}

func (app *App) Use(middlewares ...HandleFunc) {
	app.router.handlerFuncs = append(r.handlerFuncs, middlewares)
}
