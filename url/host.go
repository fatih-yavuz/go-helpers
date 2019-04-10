package url

import "net/url"

func Host(u string) (string, error) {
	parsedUrl, err := url.ParseRequestURI(u);
	if err != nil {
		return "", err
	}
	return parsedUrl.Host, nil
}
