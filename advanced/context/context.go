package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	todoContext := context.TODO()
	backgroundContext := context.Background()

	ctx := context.WithValue(todoContext, "name", "golang")

	fmt.Println(todoContext)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	fmt.Println("----------------------------------------------------------------")

	ctx = context.WithValue(backgroundContext, "City", "Raipur")

	fmt.Println(backgroundContext)
	fmt.Println(ctx)
	fmt.Println(ctx.Value("City"))

	fmt.Println("----------------------------------------------------------------")

	ctx = context.TODO()

	result := oddOrEven(ctx, 200)
	fmt.Println(result)

	ctx = context.Background()
	//if from creation of this, if 0.5 secs pass it will result in timeout
	ctx, cancel1 := context.WithTimeout(ctx, time.Second/4)
	defer cancel1()

	result = oddOrEven(ctx, 245)
	fmt.Println(result)

	// time slept
	time.Sleep((time.Second / 2) + (2 * time.Millisecond))
	fmt.Println((time.Second / 2) + (2 * time.Millisecond))

	result = oddOrEven(ctx, 300)
	fmt.Println(result)

	fmt.Println("-----------------------------------------------")

	// context.Background() is he root context from where all other context can be derived
	rootContext := context.Background()

	worker, cancel2 := context.WithTimeout(rootContext, 2* time.Second)

	//after 2 seconds it will send a cancellation signal
	//but it will still retain its values and values can also be added
	id := context.WithValue(worker,"request","Golang")

	defer cancel2()

	go doWord(id)

	time.Sleep(3 * time.Second)

	myResult := id.Value("request")
	if myResult != nil{
		fmt.Println("result is:", myResult)
	} else{
		fmt.Println("result not found")
	}

	go doWord(id)

	id = context.WithValue(worker,"doing","Love")

	myResult = id.Value("doing")
	if myResult != nil{
		fmt.Println("result is:", myResult)
	} else{
		fmt.Println("result not found")
	}

	fmt.Println("-----------------------------------------------")

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

func doWord(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("error:", ctx.Err())
			return
		default:
			fmt.Println("working.....")
			time.Sleep(time.Second/2)
		}
	}
}
