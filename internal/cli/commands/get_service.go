package commands

import (
	"github.com/andresmeireles/speaker/internal"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

func getService[T any]() T {
	sl := servicelocator.NewServiceLocator()
	services := []servicelocator.Instantiable{}

	services = append(services, internal.Misc()...)
	services = append(services, internal.Repos()...)

	for _, service := range services {
		servicelocator.Set(sl, service)
	}

	return servicelocator.Get[T](*sl)
}
