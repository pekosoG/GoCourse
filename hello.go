package main

import  "fmt"



//This is the same type as string, but with our own name
type Salutation string

//Or we can define our Structure/Object, since Go does not have Classes as everyone else
//Also, this struct is public visible since it starts on Capital instead lowercase
type SalutationStruct struct{
	name string
	greeting string
}


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

	var phrase Salutation = "Hello World Salutation"
	fmt.Println(phrase)

	//In here we can initialize our variable sending the parameters while initializing
	var phraseStruct = SalutationStruct{"Pekosog","Hello There"}

	//There are no concepts of getters or setters, we can access vaclues as follow
	fmt.Println(phraseStruct.name)
	fmt.Println(phraseStruct.greeting)

	//We can also send the parameters in any order BUT we need to define the variable name
	phraseStruct = SalutationStruct{greeting:"Hello there",name:"PekosoG"}

	fmt.Println(phraseStruct.name)
	fmt.Println(phraseStruct.greeting)

	//Or we can initialize in blank and set the values afterwards
	phraseStruct = SalutationStruct{}
	phraseStruct.name="PekosoG"
	phraseStruct.greeting="Hello There"
}
