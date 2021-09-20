package main

import (  
    "fmt"
	"unicode/utf8"
)

func compareStrings(str1 string,str2 string){
	if str1==str2{
		fmt.Printf("\nString1= %s and string2= %s are Equal",str1,str2)
	} else{
		fmt.Printf("\nString1= %s and string2= %s are Not Equal",str1,str2)
	}
}
func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}
// func printChar(s string) {  
//     fmt.Printf("Character: ")
//     for i := 0; i < len(s); i++ {
//         fmt.Printf("%c ", s[i])
//     }
// }--> has a serious bug in terms of numbe rof bytes code occupys tehre comes the use of rune as used below

func printChars(s string) {  
    fmt.Printf("Characters: ")
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c ", runes[i])
    }
}

func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}

func main() {  
    name := "Hello World"
	fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
    fmt.Printf("\n\n")
    name = "Señor"
    fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
	for i,rune :=range name{
		fmt.Printf("\n%c starts at byte %d\n",rune,i)
	}

	//creating a string from slice of bytes
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}//hexdec
    str := string(byteSlice)
    fmt.Println("\n",str)
	byteSlice1 := []byte{67, 97, 102, 195, 169}//decimal
	str1 := string(byteSlice1)
    fmt.Println("\n",str1)
	//creating a string from slice of runes
	runeSlice1 := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}//hexdecimal
	str2 := string(runeSlice1)
    fmt.Println("\n",str2)

	//runecountinstring

	word1 := "Señor"
    fmt.Printf("String: %s\n", word1)
    fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
    fmt.Printf("Number of bytes: %d\n", len(word1))

    fmt.Printf("\n")
    word2 := "Pets"
    fmt.Printf("String: %s\n", word2)
    fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
    fmt.Printf("Number of bytes: %d\n", len(word2))

	//string compare
	string1 := "Going"
    string2 := "Going"
    compareStrings(string1, string2)

    string3 := "hello"
    string4 := "world"
    compareStrings(string3, string4)

	//str concatenationn 
	//using normal-->
	name1 := "Govind"
    last1 := "Raman"
    result := fmt.Sprintf("%s %s", name1, last1)
	fmt.Println("\n",name1," ",last1)
    fmt.Println("\n",result)

	//string are immutable but we can use rune to change
	h := "hello"
    fmt.Println(mutate([]rune(h)))
}