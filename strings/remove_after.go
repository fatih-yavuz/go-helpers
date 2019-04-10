package strings

import "strings"

func RemoveAfter(str string, after string) string {
	i := strings.Index(str, after);
	if i > 0 {
		str = str[0 : i+len(after)]
	}
	return str
}
