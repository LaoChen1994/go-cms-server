package util

import (
	"reflect"
)

func Map(data interface{}, fn interface{}) (results []interface{}) {
	vfn := reflect.ValueOf(fn).Elem()
	vdata := reflect.ValueOf(data)

	for i := 0; i < vdata.Len(); i++ {
		param := []reflect.Value{vdata.Index(i), reflect.ValueOf(i)}
		results = append(results, vfn.Call(param))
	}

	return
}

func Filter[T comparable](data []T, fn func(data T, i int) bool) (results []T) {
	vdata := reflect.ValueOf(data)
	vfn := reflect.ValueOf(fn)

	for i := 0; i < vdata.Len(); i++ {
		elem := vdata.Index(i)
		match := vfn.Call([]reflect.Value{elem, reflect.ValueOf(i)})

		if match[0].Bool() {
			results = append(results, data[i])
		}
	}

	return
}

func Find[T comparable](data []T, fn func(data T, i int) bool) (result T) {
	vdata := reflect.ValueOf(data)
	vfn := reflect.ValueOf(fn)

	for i := 0; i < vdata.Len(); i++ {
		match := vfn.Call([]reflect.Value{vdata.Index(i), reflect.ValueOf(i)})

		if match[0].Bool() {
			result = data[i]
			return
		}
	}

	return
}
