package utils

import "reflect"

func CanBeNil[R any]() bool {
	var zeroValOfR R
	rt := reflect.TypeOf(zeroValOfR)
	return rt == nil ||
		rt.Kind() == reflect.Ptr ||
		rt.Kind() == reflect.Interface ||
		rt.Kind() == reflect.Slice ||
		rt.Kind() == reflect.Map ||
		rt.Kind() == reflect.Chan ||
		rt.Kind() == reflect.Func
}
