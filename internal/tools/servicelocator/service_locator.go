package servicelocator

import "fmt"

type ServiceLocator struct {
	services map[string]any
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]any),
	}
}

func (s *ServiceLocator) Set(name string, instance any) {
	s.services[name] = instance
}

func (s *ServiceLocator) Get(name string) any {
	service := s.services[name]
	if service == nil {
		panic("service: " + name + ". not found")
	}

	return service
}

func (s *ServiceLocator) GetE(name string) (any, error) {
	service := s.services[name]
	if service == nil {
		return nil, fmt.Errorf("service not found: %s", name)
	}

	return service, nil
}

func (s ServiceLocator) GetServices() map[string]any {
	return s.services
}
