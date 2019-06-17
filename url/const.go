package url

const (
	HTTP  string = "http://"
	HTTPS string = "https://"
)

const (
	SCHEME   string = "scheme"
	USERNAME string = "username"
	PASSWORD string = "password"
	HOSTNAME string = "hostname"
	PORT     string = "port"
	PATH     string = "path"
	FILENAME string = "filename"
	QUERY    string = "query"
	FRAGMENT string = "fragment"
)

const (
	scheme   string = `?P<` + SCHEME + `>https?|ftp`
	username string = `?P<` + USERNAME + `>.*?`
	password string = `?P<` + PASSWORD + `>.*?`
	hostname string = `?P<` + HOSTNAME + `>[^:\/\s]+`
	port     string = `?P<` + PORT + `>:([^\/]*)`
	path     string = `?P<` + PATH + `>(\/\w+)*\/`
	filename string = `?P<` + FILENAME + `>[-\w.]+[^#?\s]*`
	query    string = `?P<` + QUERY + `>\?([^#]*)`
	fragment string = `?P<` + FRAGMENT + `>#(.*)`
)

const (
	pattern string = `^((%s):\/)?\/?((%s)(:(%s)|)@)?(%s)(%s)?(%s)(%s)?(%s)?(%s)?$`
)
