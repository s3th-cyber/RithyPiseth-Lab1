// Exercise 5 (binary, hex , and base64 encoding)
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main(){
	var s string
	fmt.Println("======Convert String to binary, Hex, and Base64======")
	fmt.Print("Input : ")
	fmt.Scanln(&s)

	fmt.Print("Binary : ")
	for i:=0 ; i < len(s) ; i++{
		fmt.Printf("%b ",s[i])
	}
	fmt.Print("\nHexa : ")
	hex := hex.EncodeToString([]byte(s))
	fmt.Println(hex)
	fmt.Println([]byte(s))
	fmt.Println("Base64 : ",base64.StdEncoding.EncodeToString([]byte(s)))

}