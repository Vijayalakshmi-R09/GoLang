package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{1,2,3,4,5}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
	fmt.Println("Length of num: ",len(num))
}