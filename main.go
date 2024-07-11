package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	password "github.com/juancarbajal/password-generator/pkg"
)

const MAX_SIZE = 256

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

var noNumbers bool
var noSymbols bool
var size int

var rootCmd = &cobra.Command{
	Use:   "password-generator",
	Short: "Generate a random password",
	Run: func(*cobra.Command, []string) {
		fmt.Println(password.Generate(size, !noNumbers, !noSymbols))
	},
}

// main ...
func main() {
	defer timer("main")()
	rootCmd.Flags().BoolVarP(&noSymbols, "no-symbols", "x", false, "Do not include special characters")
	rootCmd.Flags().BoolVarP(&noNumbers, "no-numbers", "n", false, "Do not include numbers")
	rootCmd.Flags().IntVarP(&size, "size", "s", MAX_SIZE, "Size of the password")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
