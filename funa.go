package main

import "fmt"

/*
 * 描述：函数壁报
 *
 *	closure: 参数   x int 
 * 		 返回值 func( int )( int, int )
 *
 *		 通过匿名函数实现
 *
 **********************************************************************/

func closure( x int ) func( int )( int, int ){

	if x == 10 {
		return func ( y int )( value , fage int ){
			value = x + y
			fage = 1
			return 
		}
	}else {
		return func ( y int )( value , fage int ){
			value = x - y
			fage = 2
			return 
		}
	}
}


func main(){
	funa := closure( 10 )
	fmt.Println( funa( 1 ) )	// 输出：11 1
	fmt.Println( funa( 2 ) )	// 输出：12 1
	fmt.Println( funa( 3 ) )	// 输出：13 1
	fmt.Println( funa( 4 ) )	// 输出：14 1

	funb := closure( 100 )
	fmt.Println( funb( 1 ) )        // 输出：99 2
	fmt.Println( funb( 2 ) )        // 输出：98 2
	fmt.Println( funb( 3 ) )        // 输出：97 2
	fmt.Println( funb( 4 ) )        // 输出：96 2

}
