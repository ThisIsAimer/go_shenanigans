package main
import "fmt"

type customError struct{
	number int
	message string
}

func (c *customError) Error() string{
	return fmt.Sprintf("Error: %d, %s", c.number,c.message)
}

func do_something() error{
	return &customError{number: 500,message: "something went wrong"}
}



func main(){
	if err := do_something(); err != nil{
		fmt.Println(err)
	}
	fmt.Println("HEY")
}