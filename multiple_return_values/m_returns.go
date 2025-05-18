package main

import (
	"errors"
	"fmt"
)

func main(){
	var q, r = devide(5,10)
	fmt.Printf("quotient: %v and remainder: %v \n", q, r)

	var result, err = compareGreater(1,1)

	if err != nil {
		fmt.Println("Error:", err)
	}else{
		println(result)
	}

}

func devide(a,b int) (quotient int, remainder int){
	quotient = a/b
	remainder = a%b
	return
}

func compareGreater (a,b int) (string, error){
	if a>b{
		return "the first value is bigger",nil
	} else if a<b {
		return "the first value is bigger",nil
	} else {
		return "", errors.New("unable to compare which is greater")
	}
}

