package main

import "fmt"

func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

func main() {  
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)
	area, _ := rectProps(10.8, 5.6)
    fmt.Printf("Area %f ", area)
}

func areacal(len int,bre int) int{
	var areas=len*bre
	return areas
}
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}
// func rectProps(length, width float64)(float64, float64) {  
//     var area = length * width
//     var perimeter = (length + width) * 2
//     return area, perimeter
// }