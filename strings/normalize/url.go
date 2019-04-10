package normalize

import (
	"net/url"
	"strings"
)

/*
	Url takes a string and returns normalized url
	Examples:
		input -> www.example.com
		ouptut -> http://www.example.com

		input -> //example.com?abc=123#anchor
		output -> https://example.com?abc=123

		input -> src="//example.com?abc=123#anchor"
		output -> https://example.com?abc=123

		input -> src='//example.com?abc=123#anchor'
		output -> https://example.com?abc=123

		input -> href="//example.com?abc=123#anchor"
		output -> https://example.com?abc=123

	Caveats:
		if the string does not contain the protocol, the assumption is http
		if url starts with double slash then the protocol is assumed to be https
*/
func Url(path string) (string, error) {
	if strings.HasPrefix(path, "src=\"", ) {
		path = strings.Replace(path, "src=\"", "", -1)
	} else if strings.HasPrefix(path, "src='", ) {
		path = strings.Replace(path, "src='", "", -1)
	} else if strings.HasPrefix(path, "href='", ) {
		path = strings.Replace(path, "href='", "", -1)
	} else if strings.HasPrefix(path, "href=\"", ) {
		path = strings.Replace(path, "href=\"", "", -1)
	}

	if strings.HasSuffix(path, "\"") || strings.HasSuffix(path, "'") {
		path = path[0 : len(path)-1]
	}

	if strings.HasPrefix(path, "//") {
		path = "https:" + path
	} else if !strings.HasPrefix(path, "http") {
		path = "http://" + path
	}

	parsed, err := url.ParseRequestURI(path);
	if err != nil {
		return "", err
	}

	result := parsed.String()

	i := strings.Index(result, "#");
	if i > 0 {
		result = result[0:i]
	}

	return result, nil
}
