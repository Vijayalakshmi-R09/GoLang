package main

import "fmt"

type detail struct{
	age int
	country string
}
func main(){
	name1:=detail{
		age: 24,
		country: "USA",
	}
	name2:=detail{
		age: 24,
		country: "UK",
	}
	name3:=detail{
		age: 26,
		country: "NZ",
	}

	name:=map[string] detail{
		"pre3":name1,
		"vj":name2,
		"ani":name3,
	}
	for name,info:=range name{
		fmt.Printf("name: %s, age: %d, country: %s\n",name,info.age,info.country)
	}

	//length map
	fmt.Println("The length of mao(name): ",len(name))

}