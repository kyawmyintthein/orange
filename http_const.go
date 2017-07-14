package orange

const (
	CharsetUTF8 string = "charset=UIF-8"
)

// MIMEType types
const (
	MIMETypeApplicationJSON                  = "application/json"
	MIMETypeApplicationJSONCharsetUTF8       = MIMETypeApplicationJSON + "; " + CharsetUTF8
	MIMETypeApplicationJavaScript            = "application/javascript"
	MIMETypeApplicationJavaScriptCharsetUTF8 = MIMETypeApplicationJavaScript + "; " + CharsetUTF8
	MIMETypeApplicationXML                   = "application/xml"
	MIMETypeApplicationXMLCharsetUTF8        = MIMETypeApplicationXML + "; " + CharsetUTF8
	MIMETypeTextXML                          = "text/xml"
	MIMETypeTextXMLCharsetUTF8               = MIMETypeTextXML + "; " + CharsetUTF8
	MIMETypeApplicationForm                  = "application/x-www-form-urlencoded"
	MIMETypeApplicationProtobuf              = "application/protobuf"
	MIMETypeApplicationMsgpack               = "application/msgpack"
	MIMETypeTextHTML                         = "text/html"
	MIMETypeTextHTMLCharsetUTF8              = MIMETypeTextHTML + "; " + CharsetUTF8
	MIMETypeTextPlain                        = "text/plain"
	MIMETypeTextPlainCharsetUTF8             = MIMETypeTextPlain + "; " + CharsetUTF8
	MIMETypeMultipartForm                    = "multipart/form-data"
	MIMETypeOctetStream                      = "application/octet-stream"
)

// Headers
const (
	HeaderAccept              = "Accept"
	HeaderAcceptLanguage      = "Accept-Language"
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

const (
	ProtocolHttp   = "http"
	ProtocolHttps  = "https"
)
