package main

import (
	"errors"
	"fmt"
)

type customError struct{
	number int
	message string
	er error
}

func (c *customError) Error() string{
	return fmt.Sprintf("Error: %d, %s, %v", c.number,c.message,c.er)
}

func errFunc() error{
	return errors.New("new error")
}

func do_something() error{
	var err = errFunc()
	return &customError{number: 500,message: "something went wrong",er: err}
}



func main(){
	if err := do_something(); err != nil{
		fmt.Println(err)
	}
	fmt.Println("HEY")
}