package orange

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"sync"
)

const(
	defaultAppName = "Orange"
	defaultConfig     = "application.yaml"
	defaultConfigPath = "config"
)
type HandlerFunc func(ctx *Context)
type config *viper.Viper
type App struct {
	Name string
	*Namespace
	rootDir string
	Config  *config
	ConfigPath string
	router  *httprouter.Router
	render  Render
	pool    sync.Pool
}

func New() *App {
	var app *App
	app = new(App)
	app.Name = defaultAppName
	app.ConfigPath = defaultConfigPath
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

func (app *App) loadConfig(path string) {
	var config *config
	replacer := strings.NewReplacer(".", "_")
	config = viper.New()
	config.SetEnvKeyReplacer(replacer)
	config.SetEnvPrefix(prefix)
	config.AutomaticEnv()
	config.SetConfigName(name)
	config.AddConfigPath(path)
	config.SetConfigType(filetype)
	err := config.ReadInConfig()
	if err != nil {
	}
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
	})
	app.Config = config
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
