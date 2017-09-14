package main

import "fmt"

func main(){

	LABEL1:
	for{
		for i := 0; i< 10; i++ {

			if( i == 5 ){
				break LABEL1
			}
		}
	}
	fmt.Println( "Over" )
}
