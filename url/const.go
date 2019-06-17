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
	scheme   string = `?P<scheme>https?|ftp`
	username string = `?P<username>.*?`
	password string = `?P<password>.*?`
	hostname string = `?P<hostname>[^:\/\s]+`
	port     string = `?P<port>:([^\/]*)`
	path     string = `?P<path>(\/\w+)*\/`
	filename string = `?P<filename>[-\w.]+[^#?\s]*`
	query    string = `?P<query>\?([^#]*)`
	fragment string = `?P<fragment>#(.*)`
)

const (
	pattern string = `^((%s):\/)?\/?((%s)(:(%s)|)@)?(%s)(%s)?(%s)(%s)?(%s)?(%s)?$`
)
