package main

import (
	"errors"
	"fmt"
	"math"
)


func sqrt(x float64) (float64,error){
	if x<0 {
		return 0, errors.New("Math: sqrt of negetive number")
	}else{
		var root = math.Sqrt(x)
		return root, nil
	}
}  

func main() {
	var result, err = sqrt(-10)
	if err != nil{
		fmt.Println(fmt.Errorf("warning is %w", err))
	}else{
		fmt.Println(result)
	}
	fmt.Println("hello")

	fmt.Println(eProcess("my custom error"))
	fmt.Println(readData())

	
}

type myError struct{
	message string
}


// error is an interface with the Error method that returns a string.
//so, we can make myError type error with this. error that returns
// myError will activate this function
func (m *myError) Error() string{
	return fmt.Sprintln("Error:", m.message)
}

func eProcess(m string ) error{
	return &myError{message: m}
	// return &myError{message: "my error"}
}
//------------------------------------------------


func readData() error{
	err := readConfug()
	if err!= nil{
		return	fmt.Errorf("error is: %w", err)//%w is for errors
	}else{
		return nil
	}
}

func readConfug() error{
	return errors.New("read config error")
}