package main

import (
	"fmt"
	password "github.com/juancarbajal/password-generator/pkg"
)
const MAX_SIZE=128
// main ...
func main()  {
	fmt.Println(password.Generate(MAX_SIZE, true, true))
}
