// Exercise 3 ( Bitwise and Assignment opr )
package main

import "fmt"


// myOR
func oR(a,b int) int {
	return a|b
}
//myAND
func and(a,b int) int{
	return a&b
}
// myNotOR
func nor(a,b int) int{
	return a^b
}
// mynot
func not(x int) int {
	return  ^x
}
// myleftShift
func leftShift(x int) int {
	return x<<2
}

// myrightShift
func rightShift(x int) int {
	return x >> 1
}

func main(){
	var a,b int;

	fmt.Print("Input 2 number : ")
	fmt.Scan(&a,&b)

	fmt.Printf("a & b : %b\n",and(a,b))
	fmt.Printf("a | b : %b\n",oR(a,b))
	fmt.Printf("a NOR b : %b\n",nor(a,b))
	fmt.Printf("Not a : %b\nNot b : %b\n",not(a),not(b))
	fmt.Printf("2 left shift of a : %b\n3 2 left shift of b : %b\n",leftShift(a),leftShift(b))
	fmt.Printf("1 rigth shift of a : %b\n1 right shift of b : %b\n",rightShift(a),rightShift(b))

	// assignment operator

	x :=15
	fmt.Println("x: ",x+5)
	fmt.Println("x: ",x-10)
	fmt.Println("x: ",x*5)
	fmt.Println("x: ",x/2)
	y :=20
	fmt.Println("y: ",y%2)

}