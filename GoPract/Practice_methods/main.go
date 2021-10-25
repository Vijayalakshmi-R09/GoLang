package main

import (  
    "fmt"
    "math"
)
type Employee struct{
	name string
	age int
}
type Sqr struct{
	side float64
}
type address struct {  
    city  string
    state string
}
type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b

func (a address) fullAddress() {  
    fmt.Printf("\nFull address: %s, %s", a.city, a.state)
}

type person struct {  
    firstName string
    lastName  string
    address
}

func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Method with pointer receiver  
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}
//displaydeta method has employyee struct as receiver dtype
func (e Employee) displayDeta(){
	fmt.Printf("\nName is %s and age is %d ",e.name,e.age)
}

func (s Sqr) displayDeta(){
	fmt.Printf("\nThe sq.root of %f = %f",s.side,math.Sqrt(s.side))
}
func main(){
	e1:=Employee{
		name: "Vj",
		age: 24,
	}
	s1:=Sqr{
		side:52.33,
	}
	e1.displayDeta()//calling method
	s1.displayDeta()

	fmt.Printf("Employee name before change: %s", e1.name)
    e1.changeName("Dj")
    fmt.Printf("\nEmployee name after change: %s", e1.name)

    fmt.Printf("\n\nEmployee age before change: %d", e1.age)
    e1.changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e1.age)

	p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }

    p.fullAddress()
	num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
