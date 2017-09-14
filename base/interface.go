package main

import "fmt"

type USB interface{
	Name() string
	Connect			// 嵌入接口
}

type Connect interface{
	Connect()
}

type PhoneConnect struct{
	name string
}

func ( pc PhoneConnect ) Name() string {
	return pc.name
}

func ( pc PhoneConnect ) Connect() {
	fmt.Println( "Connect:", pc.name )
}

func main(){
	var usb USB
	a := 100
	usb = PhoneConnect{"PhoneUSB"}
	usb.Connect()
	Disconnect( usb )
	Disconnectnull( usb )
	Disconnectnull( a )
}

func Disconnect( usb USB ){

	if pc, ok := usb.( PhoneConnect ); ok {
		fmt.Println("Disconnect:", pc.name )
		return
	}
	fmt.Println("unknown decive")
}

// 空接口的用法
func Disconnectnull( usb interface{} ){

	switch v := usb.(type){
	case PhoneConnect:
		fmt.Println("Disconnect:", v.name )
	default:
		fmt.Println("unknown decive")
	}

}

