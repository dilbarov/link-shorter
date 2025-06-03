package utils

import "reflect"

func ResolveType(val any) reflect.Type {
	if t, ok := val.(reflect.Type); ok {
		return TypeOfValue(t)
	}
	return TypeOfValue(val)
}
