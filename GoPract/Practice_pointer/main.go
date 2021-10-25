package main

import "fmt"

func main(){
	b:=210
	var a *int
	if a==nil{
		fmt.Println("the a is ",a)
		a:=&b
		fmt.Println("a after initailzation ",a)
		fmt.Printf("Type of a is %T\n",a)
		fmt.Println("address of b: ",a)
		fmt.Println("the value of b: ",*a)
		*a++
		fmt.Println("new value of b is", b)
	}

	s:=new(int)
	fmt.Printf("S value is %d, type is %T, address is %v\n", *s, s, s)
	*s=84
	fmt.Println("New s value : ",*s)
	val:=100
	fmt.Println("Val before fn call: ",val)
	v:=&val
	change(v)
	fmt.Println("Val after fn call: ",val)
	d:=h()
	fmt.Println("the val of d: ",*d)

	arr:=[3] int {90,80,70}
	fmt.Println(arr)
	modify(&arr)
	fmt.Println(arr)

	ar:=[3]int{90,80,70}
	modi(ar[:])
	fmt.Println(ar)
}

func change(val *int){
	*val=50
}

func h() *int{
	i:=5
	return &i
}

func modify(arr *[3] int){
	arr[0]=100
}

func modi(sls []int){
	sls[0]=100
}