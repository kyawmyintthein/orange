package orange

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

// MIME types
const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)

// Headers
const (
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-IP"
	HeaderXRequestID          = "X-Request-ID"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"

	// Access control
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

	// Security
	HeaderStrictTransportSecurity = "Strict-Transport-Security"
	HeaderXContentTypeOptions     = "X-Content-Type-Options"
	HeaderXXSSProtection          = "X-XSS-Protection"
	HeaderXFrameOptions           = "X-Frame-Options"
	HeaderContentSecurityPolicy   = "Content-Security-Policy"
	HeaderXCSRFToken              = "X-CSRF-Token"
)

// Error
var (
	notFoundError       = NewHttpError(http.StatusNotFound)
	internalServerError = NewHttpError(http.StatusInternalServerError)
)

// bufffer pool
//var bufPool = pool.NewBufferPool(100)

type HandlerFunc func(ctx *Context)

type App struct {
	name       string
	router     *Router
	httprouter *httprouter.Router
	pool       sync.Pool
}

// New: init new app object
func NewApp(name string) *App {
	var app *App
	app = new(App)
	app.name = name
	app.defaultPool()
	app.newRouter()
	return app
}

// defaultPool: load default pool
func (app *App) defaultPool() {
	app.pool.New = func() interface{} {
		return &Context{app: app, index: -1, data: nil}
	}
}

// newContext: init new Context for each request and response
func (app *App) newContext(rw http.ResponseWriter, req *http.Request) *Context {
	var ctx *Context
	ctx = app.pool.Get().(*Context)
	ctx.request = req
	ctx.response = &Response{app: app}
	ctx.Writer = ctx.response
	ctx.index = -1
	ctx.data = nil
	ctx.response.reset(rw)
	ctx.app = app
	return ctx
}

// newRouter: init new router
func (app *App) newRouter() {
	var (
		hrouter *httprouter.Router
	)
	hrouter = httprouter.New()
	app.router = &Router{
		handlerFuncs: nil,
		prefix:   "/",
		app:      app,
	}
	app.httprouter = hrouter
	app.handleNotFound()
	app.handlePanic()
}

// handleNotFound:  hanlder function for not found
func (app *App) handleNotFound() {
	app.httprouter.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var ctx *Context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s", ContentTypeJSON, charsetUTF8))
		ctx.response.WriteHeader(http.StatusNotFound)
		ctx.Next()
		b, _ := json.Marshal(notFoundError)
		ctx.response.Write(b)
		app.pool.Put(ctx)
	})
}

// handlePanic: handler function for panic
func (app *App) handlePanic() {
	app.httprouter.PanicHandler = func(rw http.ResponseWriter,req *http.Request,i interface {}){
		var ctx *Context
		ctx = app.newContext(rw, req)
		ctx.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s", ContentTypeJSON, charsetUTF8))
		ctx.response.WriteHeader(http.StatusInternalServerError)
		ctx.Next()
		b, _ := json.Marshal(i)
		ctx.response.Write(b)
		app.pool.Put(ctx)
	}
}

// Start: start http server
func (app *App) Start(addr string) {
	colorLog("[INFO] server start at: %s\n", addr)
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

func (app *App) Namespace(path string) *Router {
	router := Router{
		handlerFuncs: nil,
		prefix:       app.router.path(path),
		app:          app,
	}
	colorLog("[WRAN]", router.prefix)
	return &router
}

func (app *App) Use(middlewares ...HandlerFunc) {
	app.router.handlerFuncs = append(app.router.handlerFuncs, middlewares...)
}


