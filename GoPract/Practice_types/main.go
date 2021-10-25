package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a,b:=true,false
	fmt.Println(a&&b)
	fmt.Println(a||b)
	var age =30
    fmt.Println("My age is", age)
	fmt.Printf("Type of a is %T,Size of a= %d",a,unsafe.Sizeof(a))
	fmt.Printf("\nType of b is %T,Size of b= %d",b,unsafe.Sizeof(b))
	fmt.Printf("\nType of age is %T,Size of age= %d",age,unsafe.Sizeof(age))
	c, d := 5.67, 8.97
    fmt.Printf("\ntype of c %T d %T\n", c, d)//ans is dloat64 since default size is 64
    sum := c+d
    diff := c-d
    fmt.Println("sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)

	c1 := complex(5, 7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:", cadd)
    cmul := c1 * c2
    fmt.Println("product:", cmul)

	first := "Naveen"
    last := "Ramanathan"
    name := first +" "+ last
    fmt.Println("My name is",name)

	i := 55      //int
    j := 67.8    //float64
    sumw := i + int(j) //j is converted to int
    fmt.Println(sumw)

	ic := 10
    var jc float64 = float64(ic) //this statement will not work without explicit conversion
    fmt.Println("jc", jc)
}