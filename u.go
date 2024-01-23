package main

import (
	"fmt"
	"reflect"

	"github.com/andresmeireles/speaker/internal/person"
)

type Person interface {
}

func main() {
	getNameByInterface[person.PersonRepository]()
}

func getNameByInterface[T any]() {
	name := reflect.TypeOf((*T)(nil)).Elem().String()

	fmt.Println(name)
}
