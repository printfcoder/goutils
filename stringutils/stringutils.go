package stringutils

import (
	"fmt"
	"strings"
)

const (

	// SPACE is a String for a space character.
	SPACE = " "

	// EMPTY is the empty String ""
	EMPTY = ""
)

// region empty checks

// IsEmpty checks if the address of point cs is empty
func IsEmpty(cs string) bool {
	return len(cs) == 0
}

// IsNotEmpty checks if the address of point cs is empty
func IsNotEmpty(cs string) bool {
	return len(cs) > 0
}

// IsAnyEmpty checks if any of the css are empty or nil point
func IsAnyEmpty(css ...string) bool {

	if len(css) == 0 {
		return false
	}

	for _, str := range css {
		if IsEmpty(str) {
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
		if IsNotEmpty(str) {
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

// region trim

// Truncate truncates a String
func Truncate(str string, maxWidth int) string {
	ret, _ := TruncateFromWithMaxWith(str, 0, maxWidth)
	return ret
}

// TruncateFromWithMaxWith truncates a String
func TruncateFromWithMaxWith(str string, offset, maxWidth int) (ret string, err error) {
	if offset < 0 {
		return "", fmt.Errorf("offset cannot be negative")
	}

	if maxWidth < 0 {
		return "", fmt.Errorf("maxWidth cannot be negative")
	}

	if IsEmpty(str) {
		return str, nil
	}

	l := len(str)

	if offset > l {
		return EMPTY, nil
	}

	if l > maxWidth {

		var ix int
		if offset+maxWidth > l {
			ix = l
		} else {
			ix = offset + maxWidth
		}

		return SubStringBetween(str, offset, ix-1)
	}

	return SubString(str, offset)
}

// SubString returns a new string that is a substring of this string
// beginIndex means the string returned is the substring begins it and extends to the end of input``
func SubString(str string, beginIndex int) (ret string, err error) {

	if beginIndex < 0 {
		return "", fmt.Errorf("beginIndex cannot be negative")
	}

	subLen := len(str) - beginIndex

	if subLen <= 0 {
		return "", fmt.Errorf("beginIndex out of bound")
	}

	str2 := []rune(str)

	return string(str2[beginIndex:]), nil

}

// SubStringBetween returns a new string that is a substring of this string
func SubStringBetween(str string, beginIndex, endIndex int) (ret string, err error) {

	if beginIndex < 0 {
		return "", fmt.Errorf("beginIndex cannot be negative")
	}

	l := len(str)

	if endIndex > l {
		return "", fmt.Errorf("endIndex out of bound")
	}

	subLen := endIndex - beginIndex
	if subLen < 0 {
		return "", fmt.Errorf("endIndex must be bigger than beginIndex")
	}

	if beginIndex == 0 && endIndex == l {
		return str, nil
	} else {
		str2 := []rune(str)
		return string(str2[beginIndex:endIndex]), nil
	}
}

// endregion

// region

// Strip strips whitespace from the start and end of a String
func Strip(str string) string {
	return ""
}

// StripStart strips any of a set of characters from the start of a String.
// stringutils.StripStart("", *)            = ""
// stringutils.StripStart("abc", "")        = "abc"
// stringutils.StripStart("yxabc  ", "xyz") = "abc  "
func StripStart(str, stripChars string) string {
	str2 := []rune(str)
	l := len(str2)
	if l == 0 {
		return str
	}

	if IsEmpty(stripChars) {
		return str
	}

	start := 0
	var ch string
	for start != l {
		ch, _ = CharAt(str, start)
		if IndexOf(stripChars, ch) == -1 {
			break
		}
		start++
	}

	ret, _ := SubString(str, start)
	return ret

}

// StripEnd strips any of a set of characters from the end of a String
//
// An empty string ("") input returns the empty string
// stringutils.StripEnd("", *)            = ""
// stringutils.StripEnd("abc", "")        = "abc"
// stringutils.StripEnd("  abcyx", "xyz") = "  abc"
// stringutils.StripEnd("120.00", ".0")   = "12"
func StripEnd(str, stripChars string) string {
	str2 := []rune(str)
	end := len(str2)
	if end == 0 {
		return str
	}

	if IsEmpty(stripChars) {
		return str
	}

	var ch string
	for end != 0 {
		ch, _ = CharAt(str, end-1)
		if IndexOf(stripChars, ch) == -1 {
			break
		}
		end--
	}

	ret, _ := SubStringBetween(str, 0, end)
	return ret
}

// endregion

// CharAt returns the char value at the specified index.
func CharAt(str string, index int) (ret string, err error) {

	if index < 0 || index >= len(str) {
		return "", fmt.Errorf("%d index out of bound %d", index, len(str))
	}

	str2 := []rune(str)

	return string(str2[index : index+1]), nil
}

// IndexOf returns the index within this string of the first occurrence of
// the specified character.
func IndexOf(str, sub string) int {
	return strings.Index(str, sub)
}
