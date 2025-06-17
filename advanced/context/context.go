package main

import (
	"context"
	"fmt"
)

func main(){
	todoContext := context.TODO()
	contextBackground := context.Background()

	ctx := context.WithValue(todoContext,"name","golang")

	fmt.Println(todoContext)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	fmt.Println("----------------------------------------------------------------")

	ctx = context.WithValue(contextBackground,"City","Raipur")

	fmt.Println(contextBackground)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("City"))

}