package main

import (  
    "fmt"
)


func main() {  
	// var  n int
	// fmt.Println("Enter size of array")
	// fmt.Scanln(&n)
	var a [3] int
	a[0]=10
	a[1]=40
	a[2]=70
	fmt.Println(a)
//short hand declaration
	b:= [3] int {5,35,65}
	fmt.Println(b)

//not neccarry all elements must be declared in shorthand decl
	c:=[3] int{9}
	fmt.Println(c)
// len of array can be ignored an dreplace with [...]
	d:=[...] int{3,6,9,12,15,18,21,24,27,30}
	fmt.Println(d)
// arrays are value types i.e a copy of d is assigned to e and if any change made to e its wont affect d
	e:=d
	e[0]=1
	fmt.Println("e: ",e)
	fmt.Println("d: ",d)
	fmt.Println("Length of d ",len(d))
//iteration
	for i:=0;i<len(d);i++ {
		fmt.Println("index ",i,": ",d[i])
	}	
//range
	sum:=d[0]
	for _,j:=range d {
		fmt.Printf(" the element of d is %d\n",j)
		sum+=j
	}
	fmt.Println("Sum of d array: ",sum)

//multidemintin
arr := [3][2] int{
	{1,2},
	{3,6},
	{4,8}, //this comma is necessary. The compiler will complain if you omit this comma
}
printarray(arr)
var brr [3][2] int
brr[0][0] = 2
brr[0][1] = 3
brr[1][0] = 4
brr[1][1] = 5
brr[2][0] = 6
brr[2][1] = 7
fmt.Printf("\n")
printarray(brr)
}

func printarray(arr [3][2] int) {  
    for _, v1 := range arr {
        for _, v2 := range v1 {
            fmt.Printf("%d ", v2)
        }
        fmt.Printf("\n")
    }
}