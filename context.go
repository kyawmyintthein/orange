package orange

import (
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	response     ResponseWriter
	request      *http.Request
	params       httprouter.Params
	data         map[string]interface{}
	app          *App
	handlerFuncs []HandlerFunc
	index        int8
}
