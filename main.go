package main

import (
	"fmt"

	"com.github.Kinoshita0623/practice-grpc/app"
	"com.github.Kinoshita0623/practice-grpc/app/service"
)

func main() {

	fmt.Printf("hello\n")
	result := app.Sum(1, 2)
	fmt.Printf("sum:%d\n", result)
	service.Test()
}
