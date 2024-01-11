package servicelocator

type Instantiable interface {
	New(s ServiceLocator) interface{}
}
