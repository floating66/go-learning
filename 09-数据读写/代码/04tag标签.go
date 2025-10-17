package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"lisi"`
	Age int `info:"age"`
}

func findTag(str interface{}){
	t := reflect.TypeOf(str).Elem()
	for i:=0;i<t.NumField();i++{ 
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("taginfo = ",taginfo," tagdoc = ",tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}