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

func SetC(injection any, sl *ServiceLocator) {
	label := reflect.TypeOf(injection).String()
	hasParameters := reflect.TypeOf(injection).NumField() > 0

	if !hasParameters {
		instance := reflect.ValueOf(injection).Call([]reflect.Value{})[0]
		sl.Set(label, instance.Interface())

		return
	}

	params := parameters(injection)
	resolveParams := resolve(params, *sl)
	resolveInjection := reflect.ValueOf(injection).Call(resolveParams)[0].Interface()
	sl.Set(label, resolveInjection)
}

func parameters(injection any) []string {
	ref := reflect.TypeOf(injection)
	numFields := ref.NumField()
	fields := make([]string, numFields)

	for i := 0; i < numFields; i++ {
		fields[i] = ref.Field(i).Type.String()
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
