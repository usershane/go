package main

import (
	"fmt"
	"reflect"
)

type User struct{
	Id	int
	Name	string
	Age	int
}

type Manager struct{
	User
	title	string
}


func main(){

	mana := Manager{ User:User{1,"shane",30}, title:"123456" }

	t := reflect.TypeOf( mana )

	fmt.Printf( "%#v\n", t.FieldByIndex(( []int{0.0})) )
}

