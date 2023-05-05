// 包声明语句
package main

import "fmt"

func main()  {
	test()
}

func test()  {
	const test = 1
	var str string 
	var a,b,c int 
	var e,f = 1,false 
	p:=&f
	*p = true
	i,j :=0,1
	fmt.Println("test",test)
	fmt.Println("str",str)
	fmt.Println("data:",a,b,c)
	fmt.Println("stat:",e,f)
	fmt.Println("stat:",i,j)
}