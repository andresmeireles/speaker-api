package servicelocator

import (
	"reflect"
)

func Get[T any](service ServiceLocator) T {
	var instance T

	serviceName := reflect.TypeOf(instance).String()
	s, ok := service.Get(serviceName).(T)

	if !ok {
		panic("service not found " + serviceName)
	}

	return s
}
