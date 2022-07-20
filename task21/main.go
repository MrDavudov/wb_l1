package main

import (
	"fmt"
	"strings"
)

type StrList struct {
	rows []string
}

func (s StrList) getString() string {
	return strings.Join(s.rows, "\n")
}

func (s *StrList) add(value string) {
	s.rows = append(s.rows, value)
}

type TextAdapter struct {
	StrList
}

func (t TextAdapter) getText() string {
	return t.getString()
}

func getTextAdapter() TextAdapter {
	adapter := TextAdapter{}
	adapter.add("line 1")
	adapter.add("line 2")
	return adapter
}

func main() {
	adapter := getTextAdapter()
	text := adapter.getText()
	
	fmt.Println(text)
}
