package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)

}

// func main() {
// 	files, _ := filepath.Glob("[")
// 	fmt.Println("matched files", files)
// }---having matched files no error
