package main

import "fmt"

type Salute struct{
	name, greeting string
}

func Greet(salute Salute){
	fmt.Println(CreateMessage(salute.name,salute.greeting))

	message, alternateMessage := CreateDoubleMessage(salute.name, salute.greeting)

	fmt.Println(message) //If we do not use the first value returned, the compile process will fails
	fmt.Println(alternateMessage)

	_, alternateMessageOnly := CreateDoubleMessage(salute.name, salute.greeting) //We can throw one value away using an underscore
	fmt.Println(alternateMessageOnly)

	_, alternateNamedMessage:= CreateDoubleNamedMessage(salute.name,salute.greeting)
	fmt.Println(message)
	fmt.Println(alternateNamedMessage)

	_, alternateVariadicMessage:= CreateVariadicMessage(salute.name, salute.greeting, "What's up")
	fmt.Println(alternateVariadicMessage)
}

func CreateMessage(name, greeting string) string{
	return greeting+ " "+name
}

func CreateDoubleMessage(name,greeting string) (string, string){ //This is how we return double values
	return greeting+" "+greeting, "HEY "+name+"!!!!"
}

func CreateDoubleNamedMessage(name, greeting string)(message, alternate string){
	message= greeting+" "+name
	alternate= "HEY "+name+"!!!!"
	return
}

func CreateVariadicMessage(name string, greeting ... string)(message string, alternative string){
	fmt.Println(len(greeting)) // We can validate the length of the array of parameters.
	message= greeting[0]+" "+name
	alternative= greeting[1]+" "+name
	return
}


func main(){
	var salute = Salute{"Pekoso","Hello there"}
	Greet(salute)
}