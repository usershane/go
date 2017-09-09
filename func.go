package main

import "fmt"

var (
	ga = Func1
)

/*
 * 描述：函数参数的用法。
 *
 *	以下声明方式 ( a, b , c int ) a,b,c 都是整形
 *
 *************************************************/
func FuncPara1( a, b , c int )int{
	return a + b + c
}

/*
 * 描述：函数返回多个值的用法。
 * 
 * 	返回 c, d 两个整数. 用法 c, d = a, b 在函数声明是已经存在
 *
 *************************************************/
func FuncRetu( a, b int )( c, d int ) {
	c, d = a, b
	return c, d
}

/*
 * 描述：函数参数的用法。
 *
 * 	函数可变长参数的用法。
 *
 *************************************************/
func FuncPara2( strMax string, a ...int ){
	fmt.Println( strMax )
	fmt.Println( a )
}

func Func1(){
	fmt.Println( "Func = 1" )
}

func Func2(){
	fmt.Println( "Func = 2" )
}

func Func3(){
	fmt.Println( "Func = 3" )
}

func main(){
	a := FuncPara1( 1, 2, 3 )
	fmt.Println( a )
	b, c := FuncRetu( 4, 5 )
	fmt.Println( c )
	fmt.Println( b )
	FuncPara2( "shane" ,1,2,3,4,5,6,7,8,9 )

	ac :=  Func2
	ac()

	ga()

}
