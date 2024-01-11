package testdata

import (
	"github.com/andresmeireles/speaker/internal"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

func GetServiceLocator() *servicelocator.ServiceLocator {
	services := internal.Services()
	sl := servicelocator.NewServiceLocator()

	for _, service := range services {
		servicelocator.Set(sl, service)
	}

	return sl
}

func GetService[T any]() T {
	sl := GetServiceLocator()

	return servicelocator.Get[T](*sl)
}
