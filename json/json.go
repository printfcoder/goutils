package json

import (
	"encoding/json"
	"reflect"
)

// Copy makes a new object which has same field key and value from 'src'
func Copy(dst interface{}, src interface{}) (err error) {

	if dst == nil || (reflect.ValueOf(dst).Kind() == reflect.Ptr && reflect.ValueOf(dst).IsNil()) {
		dst = reflect.New(reflect.TypeOf(dst))
	}
	if src == nil || (reflect.ValueOf(src).Kind() == reflect.Ptr && reflect.ValueOf(src).IsNil()) {
		src = reflect.New(reflect.TypeOf(src))
	}

	bytes, err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return
	}
	return nil
}
