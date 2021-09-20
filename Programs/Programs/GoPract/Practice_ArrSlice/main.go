package main

import (  
    "fmt"
)

func main() {  
   	a:=[4] int {10,20,30,40}
	var b[]int =a[1:3]
    fmt.Println(b)
	c := []int{6, 7, 8}
	fmt.Println(c)
	darr := [...]int{5, 8, 9, 8, 10, 7, 6, 9, 5}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 

	n:=[...] int {1,2,3,4,5,6}
	n1:=n[:]
	n2:=n[:]
	fmt.Println("Array before change n: ",n)
	n1[0]=0
	fmt.Println("Array after change to n1",n)
	n2[1]=1
	fmt.Println("Array after change to n2",n)
	fmt.Println("length of n1 ",len(n1),"capacity of n1",cap(n1))

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

	//make func

	i:=make([] int,5,5)
	fmt.Println(i)//default value 00000

	name:=[] string{"vg","JJ","Pre3"}
    fmt.Println("name:", name, "has old length", len(name), "and capacity", cap(name)) 
    name = append(name, "tae")
    fmt.Println("name:", name, "has new length", len(name), "and capacity", cap(name)) 

	//nil
	var bhu []string
	if bhu==nil{
		fmt.Println("nil")
		bhu=append(bhu, "dnsk","sdhsjih","kdhsjid")
		fmt.Println(bhu)
	}
	veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)

	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside

	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
		}
    for _,v1:=range pls{
		for _,v2:=range v1{
			fmt.Printf("%s ",v2)
		}
		fmt.Printf("\n")
	}

	numbs:=copyslice()
	fmt.Println(numbs)
}
func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }
}

func copyslice() [] int{
	numb:=[] int {1,2,3,4,5,6,7,8,9}
	numb1:=numb[:len(numb)-2]
	cpy:=make([] int,len(numb1))
	copy(cpy,numb1)
	return cpy
}