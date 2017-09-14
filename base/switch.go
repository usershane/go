package main

import "fmt"

func main(){

	a := 1

	// switch 用法1
	switch a{

	case 0:
		fmt.Println( a )
	case 1:
		fmt.Println( a )
	}

	// switch 用法2
	switch {
	case a >= 0:
		fmt.Println( a )
		fallthrough

	case a >= 10:
		fmt.Println( a )
	default :
		fmt.Println( a )
	}
}
