package http

var (
	HTTPHeadersRequest = (&Set{}).Assign(
		aim,
		accept,
		acceptCharset,
		acceptDatetime,
		acceptEncoding,
		acceptLanguage,
		accessControlRequestMethod,
		accessControlRequestHeaders,
		authorization,
		cacheControl,
		connection,
		contentLength,
		contentType,
		cookie,
		date,
		dnt,
		expect,
		forwarded,
		from,
		host,
		ifMatch,
		ifModifiedSince,
		ifNoneMatch,
		ifRange,
		ifUnmodifiedSince,
		maxForwards,
		origin,
		pragma,
		rnge,
		referer,
		te,
		userAgent,
		upgrade,
		upgradeInsecureRequests)
)

var (
	HTTPHeadersResponse = (&Set{}).Assign(
		age,
		cacheControl,
		connection,
		contentEncoding,
		contentLanguage,
		contentLength,
		contentLocation,
		date,
		eTag,
		expires,
		location,
		pragma,
		setCookie,
		tk,
		upgrade,
		vary,
		via,
		warning)
)

var (
	HTTPHeaders = &Headers{
		accept:                  "application/json, text/json",
		acceptCharset:           "utf-8, iso-8859-1;q=0.5",
		acceptEncoding:          "gzip, deflate, br",
		cacheControl:            "max-age=0",
		connection:              "upgrade",
		contentLanguage:         "en-US",
		cookie:                  "app=golang;",
		dnt:                     "1",
		host:                    "www.amiiboapi.com",
		maxForwards:             "1",
		upgradeInsecureRequests: "1"}
)
