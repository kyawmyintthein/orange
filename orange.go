package orange

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"sync"
)

var(
	notFoundError = NewHttpError(http.StatusNotFound)
	internalServerError = NewHttpError(http.InternalServerError)
)

var bufPool = pool.NewBufferPool(100)
type HandlerFunc func(ctx *Context)
type config *viper.Viper
type App struct {
	*Router
	rootDir string
	router  *httprouter.Router
	pool    sync.Pool
}

func New() *App {
	var app *App
	app = new(App)
	app.defaultPool()
	app.newRouter()
	return app
}

func (app *App) defaultPool() {
	app.pool{
		New: func() interface{} {
			return app.newContext()
		},
	}
}

func (app *App) newContext(w http.ResponseWriter, r *http.Request) *Context {
	var ctx *context
	ctx = new(context)
	ctx := app.pool.Get().(*context)
	ctx.Request = response: &Response{app: app, Writer: w},
	ctx.index = -1
	ctx.data = nil
	ctx.app = app
	return ctx
}

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

// handleNotFound
func (app *App) handleNotFound() {
	app.router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx Context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s",ContentTypeJSON, UTF8)
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

// handlePanic
func (app *App) handlePanic() {
	app.router.PanicHandler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx Context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s",ContentTypeJSON, UTF8)
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

func (app *App) ServceHttp(res http.ResponseWriter, req *http.Request) {
	app.router.ServeHttp(res, req)
}

func (app *App) Start(addr stirng) {
	ColorLog("[INFO] server start at: %s\n", addr)
	if err := http.ListenAndServe(address, app); err != nil {
		panic(err)
	}
}

func (app *App) StartTLS(addr string, cert string, key string) {
	if err := http.ListenAndServeTLS(addr, cert, key, app); err != nil {
		panic(err)
	}
}

