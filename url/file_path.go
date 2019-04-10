package url

import (
	"net/url"
	"strings"
)

func FilePath(u string) (string, error) {
	parsedUrl, err := url.ParseRequestURI(u)
	if err != nil {
		return "", err
	}

	res := parsedUrl.Host + parsedUrl.Path

	if strings.HasSuffix(res, "/") {
		res = res[0 : len(res)-1]
	}

	return res, nil
}
