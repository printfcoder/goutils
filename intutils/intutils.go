package intutils

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// RandInt rand num between min and max
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

// RandIntWithDateAsPrefix rand num between min and max which take date as prefix
func RandIntWithDateAsPrefix(min, max int) uint64 {
	suf := RandInt(min, max)
	return uint64(GetNowDate()*int(math.Pow(10, float64(len(strconv.Itoa(suf))))) + suf)
}

// GetNowDate return number formatted `yyyyMMdd` of now
// eg. 2018-10-28 00.00.01 -> 20181028
func GetNowDate() int {

	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()

	if month > 9 {
		year *= int(math.Pow(10, float64(len(strconv.Itoa(year)))))
	} else {
		year *= int(math.Pow(10, float64(len(strconv.Itoa(year)))+1))
	}

	return year + month*100 + day
}

// IntArrayToStringArray converts int an array to new string array
func IntArrayToStringArray(in []int) []string {
	ret := make([]string, 0, len(in))
	for _, v := range in {
		ret = append(ret, strconv.Itoa(v))
	}
	return ret
}

// ToFloat32 converts a int num to a float32 num
func ToFloat32(in int) float32 {
	return float32(in)
}

// ToFloat64 converts a int num to a float64 num
func ToFloat64(in int) float64 {
	return float64(in)
}
