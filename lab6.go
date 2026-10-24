package main

import "fmt"
import "os"
import "bufio"

func xorEncrypt(text string, key byte) string{
	value_of_text := make([]byte,len(text))

	
	// make the text xOR with the key then store it in a list of byte then will convert it back to string
	for i:=0 ; i<len(text) ; i++{
		value_of_text[i] = text[i] ^ key
	}
	// final output
	return string(value_of_text)
}

func xorDecrypt(text string, key byte) string{
	value_of_text := make([]byte,len(text))

	
	// make the xOR to the not xOR
	for i:=0 ; i<len(text) ; i++{
		value_of_text[i] = text[i] ^ key
	}
	// final output
	return string(value_of_text)
}

func readwithSpace(message string) string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()

	return message
}

func main(){

	var input int

	var message string
	var key byte
	


	for {
    fmt.Println("====== Make your message secret here ====== ")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice : ")
	fmt.Scan(&input)

	switch input{
	case 1:
		fmt.Print("Input your message : ")
		message = readwithSpace(message)
		fmt.Print("Input a secret key between 0-255 :")
		fmt.Scan(&key)
		fmt.Println("Original: ",message)
	    fmt.Println("Encrypted: ",xorEncrypt(message,key))
		fmt.Println("======================= Done ============================")
		return
	case 2:
		fmt.Print("Input your message : ")
		message = readwithSpace(message)
		fmt.Print("Input a secret key between 0-255 :")
		fmt.Scan(&key)
		fmt.Println("Given: ",message)
	    fmt.Println("Decrypt: ",xorDecrypt(message,key))
		fmt.Println("======================= Done ============================")
		return
	case 3:
		fmt.Println("======================= Terminate the program ============================")
		return
	default:
		fmt.Println("Invalid Choice, Try again.")
	}
	}

}
