package plate

import (
	"regexp"
)

func PlateValidate(plate string) (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z]{3}[0-9][A-Za-z0-9][0-9]{2}$", plate)
	return match, err
}
