package main

import(
	"fmt"
)

func mapoper(){
	var se = make( []map[int]string, 5 )
	for i := 

}

func main(){
	var a map[int]string = make(map[int]string)	// map 的创建方式1
	b		    := make(map[string]int)	// map 的创建方式2
	//c		    := map[int]string		// map 的创建方式3
	fmt.Println(a)
	b["shane"] = 100
	fmt.Println(b)
	v,f := b["shane"]
	fmt.Println(f)
	fmt.Println(v)
	delete(b,"shane")
	fmt.Println(b)
	mapoper()
}
