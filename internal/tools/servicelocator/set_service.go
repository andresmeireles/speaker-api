package servicelocator

import (
	"fmt"
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

func SetC(injection any, sl *ServiceLocator) {
	hasParameters := reflect.TypeOf(injection).NumIn() > 0

	if !hasParameters {
		instance := reflect.ValueOf(injection).Call([]reflect.Value{})[0]
		label := reflect.TypeOf(instance.Interface()).String()
		fmt.Println(label)
		sl.Set(label, instance.Interface())

		return
	}

	params := parameters(injection)
	resolveParams := resolve(params, *sl)
	resolveInjection := reflect.ValueOf(injection).Call(resolveParams)[0].Interface()
	label := reflect.TypeOf(resolveInjection).String()
	fmt.Println(label)

	sl.Set(label, resolveInjection)
}

func parameters(injection any) []string {
	ref := reflect.TypeOf(injection)
	numFields := ref.NumIn()
	fields := make([]string, numFields)

	for i := 0; i < numFields; i++ {
		fields[i] = ref.In(i).String()
	}

	return fields
}

func resolve(params []string, sl ServiceLocator) []reflect.Value {
	values := make([]reflect.Value, len(params))

	for i, param := range params {
		values[i] = reflect.ValueOf(sl.Get(param))
	}

	return values
}
