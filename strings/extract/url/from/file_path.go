package from

import (
	"errors"
	"regexp"
)

/*
FilePath receives a string that is a file path and captures url in it
Example:
	input -> fast_upload_booking.1_20180716_084221/_https___www.booking.com_hotel_br_hotur.en-gb.html_label=gen173nr-1BCAsoIEIFaG90dXJIM1gEaIECiAEBmAEuu_profile.html
	output -> https://www.booking.com

Caveats:
	-Currently only returns host part and only works on top level domain of com
*/
func FilePath(path string) (string, error) {
	pattern := `((www\.)?[a-zA-Z]*\.com)`
	r, err := regexp.Compile(pattern)

	if err != nil {
		return "", err
	}

	res := r.FindSubmatch([]byte(path))

	if len(res) == 0 {
		return "", errors.New("No match")
	}

	return string(res[0]), nil
}
