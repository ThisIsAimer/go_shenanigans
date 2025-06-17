package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	todoContext := context.TODO()
	contextBackground := context.Background()

	ctx := context.WithValue(todoContext, "name", "golang")

	fmt.Println(todoContext)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	fmt.Println("----------------------------------------------------------------")

	ctx = context.WithValue(contextBackground, "City", "Raipur")

	fmt.Println(contextBackground)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("City"))

	fmt.Println("----------------------------------------------------------------")

	ctx = context.TODO()

	result := oddOrEven(ctx, 200)
	fmt.Println(result)

	ctx = context.Background()
	//if from creation of this, if 0.5 secs pass it will result in timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second/2)
	defer cancel()

	result = oddOrEven(ctx, 245)
	fmt.Println(result)

	// time slept
	time.Sleep(time.Second + (2 * time.Millisecond))
	fmt.Println(time.Second + (2 * time.Millisecond))

	result = oddOrEven(ctx, 300)
	fmt.Println(result)

}

func oddOrEven(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return fmt.Sprint("Function timedout")
	default:
		if num%2 == 0 {
			return fmt.Sprintf("the number %d is even", num)
		} else {
			return fmt.Sprintf("the number %d is odd", num)
		}

	}
}
