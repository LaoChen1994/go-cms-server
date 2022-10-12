package util

import (
	"reflect"
)

func Map[T comparable, K comparable](data []K, fn func(i int32, v K) T) (results []T) {
	for i := 0; i < len(data); i++ {
		results = append(results, fn(int32(i), data[i]))
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
