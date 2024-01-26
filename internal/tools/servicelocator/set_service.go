// TODO: clean this up
package servicelocator

import (
	"reflect"
)

const MAX_RECURSION = 1000

type Dependency struct {
	Name           string
	refType        reflect.Type
	Implementation any
}

func AddDependency[T any](impementation any) Dependency {
	name := reflect.TypeOf((*T)(nil)).Elem()

	return Dependency{
		Name:           name.String(),
		refType:        name,
		Implementation: impementation,
	}
}

func Set(sl ServiceLocator, dependencies []Dependency) {
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
			err := resolve(sl, deps[0])
			if err != nil {
				panic("error when create last service " + err.Error())
			}

			return
		}

		if index > maxIndex-1 {
			index = 0
		}

		err := resolve(sl, deps[index])
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

func resolve(sl ServiceLocator, dep Dependency) error {
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

	imp := reflect.ValueOf(dep.Implementation).Call(values)[0]
	isInterface := dep.refType.Kind() == reflect.Interface

	if isInterface && !imp.Type().Implements(dep.refType) {
		panic("service: " + imp.String() + " does not implement " + dep.refType.String())
	}

	sl.Set(dep.Name, imp.Interface())

	return nil
}
