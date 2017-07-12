package orange

import (
	"github.com/julienschmidt/httprouter"
)

type ContentType string
type Charset string

const (
	contentType    = "Content-Type"
	acceptLanguage = "Accept-Language"
	abortIndex     = math.MaxInt8 / 2
)

const (
	ContentTypeJSON ContentType = "applicaiton/json"
)

const (
	UTF8 ContentType = "UIF-8"
)

type Context interface {
	Request()

	// Response returns `*Response`.
	Response() *Response

	Scheme() string

	Path() string

	Param(name string) interface{}

	Params() map[string]interface{}

	QueryParam(string name) url.Values

	QueryParams() map[string]interface{}

	QueryString() string

	Form(string name) interface{}

	FormData() map[string]interface{}

	File(string name) (*multipart.FileHeader, error)

	MultipartForm() (*multipart.Form, error)

	Cookie(name string) (*http.Cooke, error)

	Cookies() []*http.Cookie

	ResponseJSON(status int, data interface{})

	ResponseJSONP(status int, callback string, data interface{})

	ResponseBlob(status int, contentType string, r io.Reader)

	App() *App
}

type context struct {
	response     ResponseWriter
	request      *http.Request
	params       httprouter.Params
	data         map[string]interface{}
	app          *App
	handlerFuncs []HandlerFunc
	index        int8
}

func (ctx *context) ResponseJSON(status int, data interface{}) {
	c.response.Header().Set(contentType, fmt.Sprintf("%s; charset=%s",ContentTypeJSON, UTF8)
	c.response.WriteHeader(status)
	if data == nil {
		return
	}
	
	buf := bufPool.Get()
	defer bufPool.Put(buf)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		panic(err)
	}

	c.response.Write(buf.Bytes())
}

//Param get param from route
func (ctx *context) Param(name string, v interface{}){
	val := c.params.ByName(name)

}

