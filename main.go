package main

import (
	"fmt"
	"reflect"
)


func main() {
	var x interface{} = map[int]int{1: 2}
	xType := reflect.TypeOf(x)
	// xValue := reflect.ValueOf(x)
	fmt.Println(xType)
}