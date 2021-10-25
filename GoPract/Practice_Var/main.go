package main

import (
	"fmt" 
    "math"
)

func main() {  
    var age int =30
    fmt.Println("My age is", age)
	age=24
	fmt.Println("My age is", age)
	age=20
	fmt.Println("My age is", age)
	var name="vj"
	fmt.Println("My name is", name)
	//multiple var declaration
	
	var width, height int
    fmt.Println("width is", width, "height is", height)
    width = 100
    height = 50
    fmt.Println("new width is", width, "new height is", height)

	var he,wi=100,50
	fmt.Println("The width is ",wi,"; Height is ",he)

	//multiple type var declaration in single 

	var(
		names="Vj"
		ph=1254
        sal=1253.01
	)
	fmt.Println(names,ph,sal)

	count,n:=10,"hmm"
	fmt.Println(count,n)

	n,b:="gmm",10
	fmt.Println(n,b)
	// n,b:="gkm",1
	// fmt.Println(n,b)

	a, d := 145.8, 543.8
    c := math.Min(a, d)
    fmt.Println("Minimum value is", c)
}