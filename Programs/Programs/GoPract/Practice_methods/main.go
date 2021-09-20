package main

import "fmt"

type Employee struct{
	name string
	age int
}

//displaydeta method has employyee struct as receiver dtype
func (e Employee) displayDeta(){
	fmt.Printf("Name is %s and age is %d ",e.name,e.age)
}
func main(){
	e1:=Employee{
		name: "Vj",
		age: 24,
	}
	e1.displayDeta()//calling method
}