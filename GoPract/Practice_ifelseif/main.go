package main

import "fmt"

func main() {
	//var a,num,c int
	// fmt.Println("Enter a: ")
	// fmt.Scanf("%d",&a)
	// fmt.Println("Enter num: ")
	// fmt.Scanf("%d",&num)
	// fmt.Println("Enter c: ")
	// fmt.Scanf("%d",&c)
	fmt.Println("Enter a: ")
	var a,num,c int
	fmt.Scanln(&a)
	fmt.Println("Enter num: ")
	fmt.Scanln(&num)
	fmt.Println("Enter c: ")
	fmt.Scanln(&c)
	if a > num && a > c {
		fmt.Printf("%d is greater than %d and %d \n", a, num, c)
	} else if num > c {
		fmt.Printf("%d is greater than %d and %d \n ", num, a, c)
	} else {
		fmt.Printf("%d is greater than %d and %d \n ", c, a, num)
	}
}
