package servicelocator

import "fmt"

type ServiceLocator interface {
	Set(name string, instance any)
	Get(name string) any
	GetE(name string) (any, error)
}

type SL struct {
	services map[string]any
}

func NewServiceLocator() *SL {
	return &SL{
		services: make(map[string]any),
	}
}

func (s *SL) Set(name string, instance any) {
	s.services[name] = instance
}

func (s *SL) Get(name string) any {
	service := s.services[name]
	if service == nil {
		panic("service: " + name + ". not found")
	}

	return service
}

func (s *SL) GetE(name string) (any, error) {
	service := s.services[name]
	if service == nil {
		return nil, fmt.Errorf("service not found: %s", name)
	}

	return service, nil
}

func (s SL) GetServices() map[string]any {
	return s.services
}
