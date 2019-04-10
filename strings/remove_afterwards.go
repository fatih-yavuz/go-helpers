package strings

import "strings"

func RemoveAfterwards(str string, after string) string {
	i := strings.Index(str, after);
	if i > 0 {
		str = str[0 : i]
	}
	return str
}
