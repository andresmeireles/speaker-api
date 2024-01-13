package servicelocator

import (
	"reflect"
)

type Dependency struct {
	Name           string
	Implementation any
}

func AddInterface(interfaceName string, implementation any) Dependency {
	return Dependency{
		Name:           interfaceName,
		Implementation: implementation,
	}
}

func AddDependency[T any](impementation any) Dependency {
	n := new(T)
	name := reflect.TypeOf(*n).String()

	return Dependency{
		Name:           name,
		Implementation: impementation,
	}
}

const MAX_RECURSION = 1000

func SetRecursive(sl *ServiceLocator, dependencies []Dependency) {
	deps := dependencies
	maxIndex := len(deps)
	index := 0
	totalIndex := 0

	for len(deps) != 0 {
		if totalIndex == MAX_RECURSION {
			panic("max recursion reached")
		}

		totalIndex++

		if maxIndex == 1 {
			resolveE(sl, deps[0])

			return
		}

		if index > maxIndex-1 {
			index = 0
		}

		err := resolveE(sl, deps[index])

		if err == nil {
			indexToRemove := index
			deps = append(deps[:indexToRemove], deps[indexToRemove+1:]...)
			index = 0
			maxIndex--

			continue
		}

		index += 1
	}
}

func resolveE(sl *ServiceLocator, dep Dependency) error {
	// fmt.Println("Add", dep.Name)
	ref := reflect.TypeOf(dep.Implementation)
	numOfParams := ref.NumIn()

	if numOfParams == 0 {
		imp := reflect.ValueOf(dep.Implementation).Call([]reflect.Value{})[0]
		sl.Set(dep.Name, imp.Interface())

		return nil
	}

	values := make([]reflect.Value, 0)

	for i := 0; i < numOfParams; i++ {
		p := ref.In(i).String()
		d, e := sl.GetE(p)

		if e != nil {
			return e
		}

		values = append(values, reflect.ValueOf(d))
	}

	imp := reflect.ValueOf(dep.Implementation).Call(values)[0].Interface()
	sl.Set(dep.Name, imp)

	return nil
}
