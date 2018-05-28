package stringutils

import "strings"

const (

	// SPACE is a String for a space character.
	SPACE = " "
)

// region empty checks

// IsEmpty checks if the address of point cs is empty
func IsEmpty(cs *string) bool {
	return cs == nil || len(*cs) == 0
}

// IsNotEmpty checks if the address of point cs is empty
func IsNotEmpty(cs *string) bool {
	return len(*cs) > 0
}

// IsAnyEmpty checks if any of the css are empty or nil point
func IsAnyEmpty(css ...string) bool {

	if len(css) == 0 {
		return false
	}

	for _, str := range css {
		if IsEmpty(&str) {
			return true
		}
	}

	return false
}

// IsNoneEmpty checks if none of the css are empty or nil point
func IsNoneEmpty(css ...string) bool {
	return !IsAnyEmpty(css...)
}

// IsAllEmpty checks if all of the css are empty or nil point.
func IsAllEmpty(css ...string) bool {

	if len(css) == 0 {
		return true
	}

	for _, str := range css {
		if IsNotEmpty(&str) {
			return false
		}
	}

	return true
}

// IsBlank checks if a cs is empty, or nil point or whitespace only.
func IsBlank(cs string) bool {

	if strLen := len(cs); strLen == 0 {
		return true
	} else {
		return len(strings.TrimSpace(cs)) != strLen
	}

}

// IsNotBlank checks if a cs is not empty, not nill and not whitespace only.
func IsNotBlank(cs string) bool {
	return !IsBlank(cs)
}

// IsAnyBlank checks if any of the css are empty or nill or whitespace only
func IsAnyBlank(css ...string) bool {
	if len(css) == 0 {
		return false
	}

	for _, str := range css {
		if IsBlank(str) {
			return true
		}
	}

	return false
}

//  endregion
