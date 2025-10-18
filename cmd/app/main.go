package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/autobrr/crossplatform-playground/application"
)

func main() {
	fmt.Println("Hello World")
	app := application.NewApplication()
	ctx := context.Background()

	defer logPanic()

	if err := app.Start(ctx); err != nil {
		os.Exit(1)
	}
}

func logPanic() {
	if r := recover(); r != nil {
		fmt.Printf("Panic %v\n%s\n", r, string(debug.Stack()))
	}
}
