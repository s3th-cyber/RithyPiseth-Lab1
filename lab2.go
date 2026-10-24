package main

import (
	"bufio"
	"fmt"
	"os"
	
)

Exercise 2 ( Logical Operator )

func main(){
	var a,b int;
	fmt.Print("Input 2 numbers : ")
	fmt.Scan(&a,&b)
	if(a>0&&b>0){
		fmt.Print("Both numbers are positive.\n")
	}else{
		fmt.Print("You must input positive numbers.\n")
	}

	if(a>b){
		fmt.Print(a," is greater than ",b)
	}else if(a!=b){
		fmt.Print("a and b both are different value")
	}else if(a==b){
		fmt.Print("both are the same number.")
	}else {
		fmt.Print("b is greater than a.")
	}
	fmt.Print("\n")
}