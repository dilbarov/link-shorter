package utils

import "reflect"

func IsEmptyResult[R any]() bool {
	var zeroValOfR R
	return reflect.TypeOf(zeroValOfR) == reflect.TypeOf(struct{}{})
}
