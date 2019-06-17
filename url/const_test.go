package url_test

const (
	URL string = "https://www.google.com/dir/1/2/search.html?arg=0-a&arg1=1-b&arg3-c#hash"
)

const (
	ERROR string = "testcase \"url_test.%s\" did not pass. test outcome \"%s\"."
)

const (
	ERROR_FUNCTION string = ERROR + " " + "function \"%s\" returned \"%v\" but testcase expected \"%v\"."
)
