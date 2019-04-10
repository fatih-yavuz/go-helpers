package url

import "net/url"

func Validate(u string) bool {
	parts, err := url.ParseRequestURI(u)
	return err == nil && parts.Host != "" && parts.Scheme != ""
}
