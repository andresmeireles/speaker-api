package servicelocator

import (
	"reflect"
)

func Get[T any](service ServiceLocator) T {
	serviceName := reflect.TypeOf((*T)(nil)).Elem().String()
	s, ok := service.Get(serviceName).(T)

	if !ok {
		panic("service not found " + serviceName)
	}

	return s
}
