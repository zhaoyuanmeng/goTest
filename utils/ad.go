package utils

import (
	"fmt"
)
func Test() {
	
	var a [3]int =  [3]int{1,2,3}
	r:= [...]int{99:2}

	mouth:=[...]string{1:"a",4:"asd",2:"b",3:"c",12:"m"}

	var runers []rune
	for _, r := range "helloworld" {
		runers = append(runers, r)
	}

	ages := map[string]int{
		"zyd":20,
		"chare":21,
	}

	// 结构体
	type Employee struct {
		ID int 
		Name string
	}

	var liming  = Employee{1,"aaa"}
	liming.ID = 2
	fmt.Println(liming)
	fmt.Println(ages["zyd"])
	fmt.Println(ages["chare"])

	fmt.Printf("%q\n",runers)

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(r[99])
	fmt.Println(mouth[1:3])
	fmt.Println(mouth)
	
	fmt.Println("arr len is",len(a))

}