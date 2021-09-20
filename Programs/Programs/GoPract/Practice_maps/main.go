package main

import "fmt"

func main() {
	emp:=make(map[string] int)
	emp["Vj"]=60000
	emp["Gj"]=80000
	emp["Pre3"]=60000
	fmt.Println(emp)

	//initailaization at the time of declaration

	detail:=map[string] int{
		"vj":24,
		"pre3":24,
		"kiran":24,
	}
	detail["ani"]=26
	fmt.Println(detail)
	name:="vj"
	age:=detail[name]
	fmt.Printf("The name: %s and age: %d",name,age)
	fmt.Println("\nThe age of lucy --> ",detail["lucy"])
	//to find key is present r not
	newname:="vj"
	value,ok:=detail[newname]
	if ok==true{
		fmt.Println("age of ",newname,"==",value)
	} else 
	{
		fmt.Println(newname,"not found")
	}
	//iterating all elements
	for key,value:= range detail{
		fmt.Printf("The key[%s] : %d \n",key,value)
	}

	//delete using delete(map,key) this func does not return any value
	delete(detail,"kiran")
	fmt.Println("details after deletion: ",detail)
	delete(detail,"kiran")
	fmt.Println("details after deletion: ",detail)
	//zero value of map is nil
	// var employeeSalary map[string]int
    // employeeSalary["steve"] = 12000 ---->throws an error 'panic: assignment to entry in nil map'

	//ref type
	modi:=detail
	modi["vj"]=23
	fmt.Println("After modification ",detail)
}