package utils

import "fmt"
func Test() {
	
	var a [3]int =  [3]int{1,2,3}
	r:= [...]int{99:2}
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(r[99])
	fmt.Println("arr len is",len(a))

}