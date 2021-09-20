package main

import (
	"fmt"
)

type Employee struct{
	fname string
	age int
	sal int
}

func main() {

    e1:=Employee{
		fname: "Vj",
		age: 24,
		sal:29785,
	}

    //creating struct without specifying field names
    e2 := Employee{"Thomas",  29, 80000}// avoid using it
	e3:=struct{
		fname string
		age int
		sal int
	}{
		fname:"pre3",
		age:24,
		sal:82000,
	}

    fmt.Println("Employee 1", e1)
    fmt.Println("Employee 2", e2)
	fmt.Println("Employee 3",e3)
	fmt.Println("Name of e1: ",e1.fname)
	fmt.Println("Name of e2: ",e2.fname)
	e3.sal=40000
	fmt.Printf("\nThe sal of e3(%s): %d\n",e3.fname,e3.sal)

	//zero value of structs
	var e4 Employee
	fmt.Println("\nEmp 4 details:\nName:", e4.fname)
    fmt.Println("Age:", e4.age)
    fmt.Println("Salary:", e4.sal)

	e5:=Employee{
		fname: "dave",
		age:25,
	}
	fmt.Println("\nEmp 5 details:\nName:", e5.fname)
    fmt.Println("Age:", e5.age)
    fmt.Println("Salary:", e5.sal)

	//craete pointer for structs
	e6:=&Employee{
		fname: "david",
		age:40,
		sal:90000,
	}
	fmt.Println("\nEmp 6 details\nName: ",(*e6).fname)
	fmt.Println("Age:", (*e6).age)
    fmt.Println("Salary:", (*e6).sal)
	fmt.Println(" name e6",e6.fname)//simpel way

	f1:=flower{
		"jasmine",
		1,
	}
	fmt.Println(f1.string)
    fmt.Println(f1.int)

	a1:=Address{
		city:"Chennai",
		state: "TN",
		detal: Employee{
			fname: "Vj",
			age: 24,
			sal:52000,
		},
	}
	fmt.Println(a1)
	
}
	//annoymous filed
type flower struct{
	string
	int
}

type Address struct { 
	detal Employee 
    city  string
    state string
}