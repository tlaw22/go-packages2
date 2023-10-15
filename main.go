package main

//
import (
	"log"

	"github.com/tlaw22/go-packages2/helpers"
)

// type go mod init github.com/tlaw22/go-packages2/helpers
const numPool = 1000

func PrintText(s string) {
	log.Println(s)
}

// create a function that used the RandomNumber function from helpers package
func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func main() {
	log.Println("Beging app...")
	// Start referencing the the package by typing the package name and then a period .
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNum = 9
	log.Println(myVar)
	PrintText("Program complete...")

	intChan := make(chan int)
	defer close(intChan) // close the channel

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}
