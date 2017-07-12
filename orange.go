package orange

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"sync"
)


type HandlerFunc func(ctx *Context)
type config *viper.Viper
type App struct {
	*Namespace
	rootDir string
	router  *httprouter.Router
	render  Render
	pool    sync.Pool
}

func New() *App {
	var app *App
	app = new(App)
	app.Name = defaultAppName
	app.defaultPool()
	app.newRouter()
	return app
}

func (app *App) ServceHttp(res http.ResponseWriter, req *http.Request) {
	app.router.ServeHttp(res, req)
}

func (app *App) Start() {

}

func (app *App) Use(middlewares ...HandleFunc) {

}

func (app *App) defaultPool() {
	app.pool{
		New: func() interface{} {
			return app.newContext()
		},
	}
}

func (app *App) newContext() *Context {
	var context *Context
	context = new(Context)
	return context
}

func (app *App) newRouter() {
	var (
		httprouter *httprouter.Router
	)
	httprouter = httprouter.New()
	app.httprouter = httprouter
	app.handleNotFound()
	app.handlePanic()
}

func (app *App) handleNotFound() {
	app.router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx Context

	})
}

func (app *App) handlePanic() {
	app.router.PanicHandler = http.HandlerFunc(func() {
	})
}
