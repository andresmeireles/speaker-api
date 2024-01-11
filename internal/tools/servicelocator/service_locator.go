package servicelocator

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
