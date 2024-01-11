package servicelocator

import (
	"reflect"
)

func Set(servicelocator *ServiceLocator, newFunc Instantiable) {
	if servicelocator == nil {
		panic("service locator is nil")
	}

	service := newFunc.New(*servicelocator)
	name := reflect.TypeOf(service).String()

	servicelocator.Set(name, service)
}
