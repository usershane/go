package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id	int
	Name    string
	Age	int
}

func ( user User ) Hello() {
	fmt.Println( "Hello ", user.Name )
}

func ( user *User )SetName( strName string ){
	user.Name = strName
}

func main(){
	user := User{1,"shane",30}
	info( user )

	user.SetName( "King" )

	fmt.Println( user )
}

func info ( o interface{} ) {
	t := reflect.TypeOf( o )		// 获取类型信息
	fmt.Println( "Type : ", t.Name() )

	f := reflect.ValueOf( o )		// 获取所包含的字段信息
	fmt.Println("Fields:")

	/*
	 * 描述：获取类型的字段信息
	 *	t.NumField() // 返回字段的数量
	 *	t.Field( i ) // 返回指定字段类型，名称
	 *	f.Field( i ).Interface() // 返回指定字段的值
	 *
	 **************************************************************/
	for i := 0; i < t.NumField(); i++ {
		k := t.Field( i )
		v := f.Field( i ).Interface()
		fmt.Printf( "%6s: %6v = %v\n", k.Name, k.Type, v  )
	}

	/*
	 * 描述：获取对象的信息
	 *	t.() // 返回方法的数量
	 *	t.Field( i ) // 返回指定字段类型，名称
	 *	f.Field( i ).Interface() // 返回指定字段的值
	 *
	 **************************************************************/
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method( i )
		fmt.Printf( "%6s: %v\n", m.Name, m.Type  )
	}
}

