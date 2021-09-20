package main

import "fmt"

func main(){
	const a=50
	fmt.Println(a)

	const (
		abc=10
		bca=20
		cab=30
	)
	fmt.Println(abc)
	fmt.Println(bca)
	fmt.Println(cab)

	//error const cannot be reassigned
	// a=20
	// fmt.Print(a)

	const n = "Sam"
    var name = n
    fmt.Printf("type %T value %v", name, name)

	// var defaultName = "Sam" //allowed
    // type myString string
    // var customName myString = "Sam" //allowed
    //customName = defaultName //not allowed mixing of types not allowed in go since its a strongly typed lang

	// const trueConst = true
    // type myBool bool
    // var defaultBool = trueConst //allowed
    // var customBool myBool = trueConst //allowed
    //defaultBool = customBool //not allowed mixing of types not allowed in go since its a strongly typed lang

	const az = 5//untyped numeric const
    var intVar int = az
    var int32Var int32 = az
    var float64Var float64 = az
    var complex64Var complex64 = az
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

	var i = 5
    var f = 5.6
    var c = 5 + 6i
    fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)

	var num=10.2/20
	fmt.Printf("\n num type %T and value %v ",num,num)
}