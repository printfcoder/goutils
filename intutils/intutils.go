package intutils

import (
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
