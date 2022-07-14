package main

import (
	"fmt"
	"reflect"
)

// Задания 14
// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа
// interface{}.

func main() {
	var ch chan int
	ch <- 1

	GetType(2.1)
	GetType("Hi")
	GetType(3)
	GetType(<-ch)
}

func GetType(v interface{}) {
	var t = reflect.TypeOf(v)
	fmt.Println(t)
}