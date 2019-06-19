package http

const (
	authenticate       string = "WWW-Authenticate"
	authorize          string = "Authorize"
	authorization      string = "Authorization"
	proxyAuthenticate  string = "Proxy-Authenticate"
	proxyAuthorization string = "Proxy-Authorization"
)

const (
	age           string = "Age"
	cacheControl  string = "Cache-Control"
	clearSiteData string = "Clear-Site-Data"
	expires       string = "Expires"
	pragma        string = "Pragma"
	warning       string = "Warning"
)

const (
	lastModified      string = "Last-Modified"
	eTag              string = "ETag"
	ifMatch           string = "If-Match"
	ifNoneMatch       string = "If-None-Match"
	ifModifiedSince   string = "If-Modified-Since"
	ifUnmodifiedSince string = "If-Unmodified-Since"
	vary              string = "Vary"
)

const (
	connection string = "Connection"
	keepAlive  string = "Keep-Alive"
)

const (
	aim string = "A-IM"
)

const (
	date string = "Date"
)

const (
	accept         string = "Accept"
	acceptCharset  string = "Accept-Charset"
	acceptDatetime string = "Accept-Datetime"
	acceptEncoding string = "Accept-Encoding"
	acceptLanguage string = "Accept-Language"
)

const (
	accessControlRequestHeaders string = "Access-Control-Request-Method"
	accessControlRequestMethod  string = "Access-Control-Request-Method"
)

const (
	expect      string = "Expect"
	maxForwards string = "Max-Forwards"
)

const (
	cookie    string = "Cookie"
	setCookie string = "Set-Cookie"
)

const (
	dnt string = "DNT"
	tk  string = "TK"
)

const (
	contentLength   string = "Content-Length"
	contentType     string = "Content-Type"
	contentEncoding string = "Content-Encoding"
	contentLanguage string = "Content-Language"
	contentLocation string = "Content-Location"
)

const (
	forwarded string = "Forwarded"
	via       string = "Via"
)

const (
	location string = "Location"
)

const (
	from           string = "From"
	host           string = "Host"
	referer        string = "Referer"
	referrerPolicy string = "Referrer-Policy"
	userAgent      string = "User-Agent"
)

const (
	upgradeInsecureRequests string = "Upgrade-Insecure-Requests"
)
