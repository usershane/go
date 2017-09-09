package main

import "fmt"

func main(){
	FunA()
	FunB()
	FunC()
}

func FunA(){
	fmt.Println(" Func - A ")
}

func FunB(){
	
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("recover ")
		}
	}()
	panic("Panic in FunB")
}

func FunC(){
	fmt.Println(" Func -C ")
}
