// Exercise 4 ( Mini calculator )
package main

import (
	"fmt"
)

// Exercise 4 ( Mini calculator )

// add function
func add(x,y int) int{
	return x+y
}

// Sub
func sub(x,y int) int{
	return x-y
}

// Mul
func mul(x,y int) int{
	return x*y
}

// Div
func div(x,y int) (int,error){
	if (y ==0) {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return x/y, nil
}
// Mod
func mod(x,y int) (int,error){
	if (y ==0) {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return x%y, nil
}

func valueOfAnB() (int,int){
	var a,b int;
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	return a,b
}

func main(){
	var input int
	for{
	fmt.Println("===== Mini Calculator =====")
	fmt.Println("1. Add.")
	fmt.Println("2. Sub.")
	fmt.Println("3. Mul.")
	fmt.Println("4. Div.")
	fmt.Println("5. Mod.")
	fmt.Println("6. Exit.")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&input)

	switch input{
	case 1:
		fmt.Println("The result a+b : ",add(valueOfAnB()))
		return
	case 2:
		fmt.Println("The result subtraction: ",sub(valueOfAnB()))
		return
	case 3:
		fmt.Println("The result mudulation : ",mul(valueOfAnB()))
		return
	case 4:
		result, err := div(valueOfAnB())
		if(err!=nil){
			fmt.Println(err)
		}else{
			fmt.Println("The result of a divided b : ", result)
		}
	case 5:
		result, err := mod(valueOfAnB())
		if(err!=nil){
			fmt.Println(err)
		}else{
			fmt.Println("The result modulation : ",result)
			return
		}
	case 6:
		return
	default:
		fmt.Println("Invalid choice, Try again!")
	}
	}
}