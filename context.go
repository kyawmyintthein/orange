package orange

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"math"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const(
	abortIndex     = math.MaxInt8 / 2
	defaultMemory  = 32 << 20 // 32 MB
)

type Context struct {
	Writer       ResponseWriter
	response     *Response
	request      *http.Request
	query        url.Values
	params       httprouter.Params
	path         string
	data         map[string]interface{}
	app          *App
	handlerFuncs []HandlerFunc
	index        int8
}

// Request: returns request
func (ctx *Context) Request() *http.Request {
	return ctx.request
}

// Response: return response
func (ctx *Context) Response() *Response {
	return ctx.response
}

// Scheme: return http protocol schame as string
func (ctx *Context) Scheme() string {
	if ctx.IsTLS() {
		return ProtocolHttps
	}
	if scheme := ctx.request.Header.Get(HeaderXForwardedProto); scheme != "" {
		return scheme
	}
	if scheme := ctx.request.Header.Get(HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}
	if ssl := ctx.request.Header.Get(HeaderXForwardedSsl); ssl == "on" {
		return "https"
	}
	if scheme := ctx.request.Header.Get(HeaderXUrlScheme); scheme != "" {
		return scheme
	}
	return "http"
}

// ClientIP: return ip address of client
func (ctx *Context) ClientIP() string {
	var (
		remoteAddress, ip string
	)
	remoteAddress = ctx.request.RemoteAddr
	if ip = ctx.request.Header.Get(HeaderXForwardedFor); ip != "" {
		remoteAddress = strings.Split(ip, ", ")[0]
	} else if ip = ctx.request.Header.Get(HeaderXRealIP); ip != "" {
		remoteAddress = ip
	} else {
		remoteAddress, _, _ = net.SplitHostPort(remoteAddress)
	}
	return remoteAddress
}

// Path: return url path
func (ctx *Context) Path() string {
	return ctx.path
}

// JSON: response json to client
func (ctx *Context) JSON(status int, data interface{}) {
	var err error
	ctx.response.Header().Set(HeaderContentType, MIMETypeApplicationJSONCharsetUTF8)
	ctx.response.WriteHeader(status)
	if data == nil {
		return
	}
	buf := bufPool.Get()
	defer bufPool.Put(buf)
	if err = json.NewEncoder(buf).Encode(data); err != nil{
		colorLog("[WARN] " + err.Error())
	}
	ctx.response.Write(buf.Bytes())
}

// JSONP: response jsonp to client
func (ctx *Context) JSONP(status int, callback string, data interface{}) {
	ctx.response.Header().Set(HeaderContentType, MIMETypeApplicationJSONCharsetUTF8)
	ctx.response.WriteHeader(status)
	if data == nil {
		return 
	}
	buf := bufPool.Get()
	defer bufPool.Put(buf)
	_ = json.NewEncoder(buf).Encode(data)
	b := []byte(callback + "(" + string(buf.Bytes()) + ")")
	ctx.response.Write(b)
}

// Bytes 
func (ctx *Context) Bytes(status int, contentType string, data []byte) {
	ctx.response.Header().Set(HeaderContentType, MIMETypeApplicationJSONCharsetUTF8)
	ctx.response.WriteHeader(status)
	ctx.response.Write(data)
}

// Param: get param from route
func (ctx *Context) Param(name string) string {
	return ctx.params.ByName(name)
}

// Param: get all params from httprouter
func (ctx *Context) Params() map[string]interface{} {
	var params = make(map[string]interface{})
	for _, param := range ctx.params {
		params[param.Key] = param.Value
	}
	return params
}

// QueryParam: get parameter by name from query string
func (ctx *Context) QueryParam(name string) string {
	return ctx.request.URL.Query().Get(name)
}

// QueryParams: get all query string parameters
func (ctx *Context) QueryParams() url.Values {
	return ctx.request.URL.Query()
}

// QueryParams: get query string
func (ctx *Context) QueryString() string {
	return ctx.request.URL.RawQuery
}

// FormValue: return form value as string
func (ctx *Context) FormValue(name string) string {
	return ctx.request.FormValue(name)
}

// FormData: return form values
func (ctx *Context) FormData() (url.Values, error) {
	var err error
	if strings.HasPrefix(ctx.request.Header.Get(HeaderContentType), MIMETypeMultipartForm) {
		if err = ctx.request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
	} else {
		if err = ctx.request.ParseForm(); err != nil {
			return nil, err
		}
	}
	return ctx.request.Form, nil
}

// File: return multipart.FileHeader, error from request
func (ctx *Context) File(name string) (*multipart.FileHeader, error) {
	_, fileheader, err := ctx.request.FormFile(name)
	return fileheader, err
}

// MultipartFrom: return multipart.Form and error from request
func (ctx *Context) MultipartForm() (*multipart.Form, error) {
	err := ctx.request.ParseMultipartForm(defaultMemory)
	return ctx.request.MultipartForm, err
}

// Cookie: return cookie via name from request header
func (ctx *Context) Cookie(name string) (*http.Cookie, error) {
	return ctx.request.Cookie(name)
}

// Cookies: return all cookies from request header
func (ctx *Context) Cookies() []*http.Cookie {
	return ctx.request.Cookies()
}

// App: return app
func (ctx *Context) App() *App {
	return ctx.app
}

func (ctx *Context) Next() {
	ctx.index++
	s := int8(len(ctx.handlerFuncs))
	for ; ctx.index < s; ctx.index++ {
		ctx.handlerFuncs[ctx.index](ctx)
	}
}

func (ctx *Context) Abort() {
	ctx.index = abortIndex
}

func (ctx *Context) IsTLS() bool {
	return false
}
