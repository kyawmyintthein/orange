package orange

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"sync"
)

const (
	appConfig = "application.yaml"
)

type HandlerFunc func(ctx *Context)
type App struct {
	*Namespace
	rootDir string
	Config  *Config
	router  *httprouter.Router
	render  Render
	pool    sync.Pool
}

func New() *App {
	var app *App
	app = new(App)
	app.defaultPool()
	app.newRouter()
	app.loadConfig(configPath)
	return app
}

func AddConfigPath(configPath string) *App {
	if configPath == "" {
		configPath = filepath.Join(app.rootDir, appConfig)
	}
	app.loadConfig(configPath)
	return app
}

func (app *App) ServceHttp(res http.ResponseWriter, req *http.Request) {
	app.router.ServeHttp(res, req)
}

func (app *App) Start() {

}

func (app *App) Use(middlewares ...HandleFunc) {

}

func (app *App) loadConfig(path string) {

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
