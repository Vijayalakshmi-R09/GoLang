package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hi welcome back i hope u have a graet year ahead...")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
