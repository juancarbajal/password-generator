package main

import (
	"fmt"
	password "github.com/juancarbajal/password-generator/pkg"
	"time"
)
const MAX_SIZE=128

func timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}
// main ...
func main()  {
	defer timer("main")()
	fmt.Println(password.Generate(MAX_SIZE, true, true))
}
