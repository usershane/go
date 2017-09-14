package main

import "fmt"

type person struct{
	Name	string
	Age	int
}

func main(){
	a := &person{
		Name: "shane",
		Age:  21,
	}
	fmt.Println( a )
	set_name( a )
	fmt.Println( a )
}

func set_name( a *person ){
	a.Name = "king"
}
