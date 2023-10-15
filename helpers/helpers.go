package helpers

import (
	"math/rand"
	"time"
)

// Create a generic type just to test the included package.
// Remember that the type name must start with a capital letter
type SomeType struct {
	TypeName string
	TypeNum  int
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixMicro())
	value := rand.Intn(n)
	return value
}
