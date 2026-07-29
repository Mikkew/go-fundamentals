package main

import "fmt"

func stringsBytes() {
	// var myString string = "Hola Mundo"
	// Devuelve el valor utf-8
	// var myByte byte = myString[0]
	// fmt.Println(myByte)

	// var s2 string = myString[3:7]
	// fmt.Println(s2)

	// var s3 string = myString[:7]
	// fmt.Println(s3)

	//Strings a slices
	var s string = "Hola Mundo"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
