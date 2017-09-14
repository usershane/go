package main

import "fmt"

func main(){
	strName := "Hello Shane!!!"
	nLen	:= len(strName)

	for i := 0; i < nLen; i++ {
		fmt.Println( strName )
	}
}
