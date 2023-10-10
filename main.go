package main

//
import (
	"log"

	"github.com/tlaw22/go-packages2/helpers"
)

func main() {
	log.Println("Beging app...")
	// Start referencing the the package by typing the package name and then a period .
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNum = 9
	log.Println(myVar)
}
