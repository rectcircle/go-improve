package basic

import (
	"fmt"
	"reflect"
)

func Reflectexperiment() {
	var a int32 = 1
	var ai interface{} = a
	at := reflect.TypeOf(a)
	apt := reflect.TypeOf(&a)
	ait := reflect.TypeOf(ai)
	fmt.Println("TypeOf(a) = ", at, "TypeOf(&a) = ", apt, "TypeOf(&ait) = ", ait)
	av := reflect.ValueOf(a)
	apv := reflect.ValueOf(&a)
	fmt.Println("ValueOf(a) = ", av, "ValueOf(&a) = ", apv)

	fmt.Println("reflect.Value.Interface = ", av.Interface().(int32))
	fmt.Println("reflect.Type == reflect.Value.Type ? ", av.Type() == at)
}