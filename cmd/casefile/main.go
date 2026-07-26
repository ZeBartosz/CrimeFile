package main

import (
	"fmt"
	"os"

	"github.com/ZeBartosz/CrimeFile/internal/app"
)

func main() {
	args := os.Args[1:]
	if err := app.App(args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
