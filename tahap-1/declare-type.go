package main

import "fmt"
import "errors"


func main() {
	var variableName1 string = "hello world"
	variableName2 := "hello world"

	fmt.Println(variableName1)
	fmt.Println(variableName2)

	// primitive: boolean, int, float, string

	// boolean
	boolVar := true

	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	// integer
	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	// float
	floatVar1 := float32(3.2)
	floatVar2 := float64(3.2)

	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T Value: %v\n", floatVar2, floatVar2)

	//bytes
	bytesVar1 := byte(255)
	fmt.Printf("Type: %T Value: %v\n", bytesVar1, bytesVar1)

	bytesVar2 := []byte("hello world")
	fmt.Printf("Type: %T Value: %v\n", bytesVar2, bytesVar2)

	// rune
	runeVar := '😊'
	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	// complex
	complexVar := -7 + 3i
	fmt.Printf("Type: %T Value: %v\n", complexVar, complexVar)

	// error
	errVar := errors.New("error detected")
	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)

	// interface
	var myInterfaceVar interface {}

	myInterfaceVar = 5
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)
	myInterfaceVar = "Hello World"
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)
}

type MethodList interface {
	Myfunction()
	Myfunction2(int) int
}