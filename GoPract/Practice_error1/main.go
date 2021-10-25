package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if err != nil { //patherror defined
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
	fmt.Println(f.Name(), "opened successfully")
}
