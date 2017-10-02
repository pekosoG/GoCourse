package main

import  "fmt"

func main()  {
	//Go Declares types after the variable
	var message string

	message = "Hello world"

	fmt.Println(message)

	//Multiple variables at the same time and values as well
	var a,b,c int = 3,4,5

	fmt.Println(a,b,c)

	var messageStr = "Hello Go world"

	fmt.Println(messageStr)

	messagePoints := "Hello Go World DoublePoints"
	x,y,z:= 2, false, "zZz"

	fmt.Println(messagePoints,x,y,z)

	//Pointers behave as usual
	var greeting *string = &message

	fmt.Println(message,greeting) //Prints the address
	fmt.Println(message,*greeting) //prints the value

	*greeting = "Hi" //We modify the value of the actual message variable

	fmt.Println(message,*greeting) //prints the new value

}
