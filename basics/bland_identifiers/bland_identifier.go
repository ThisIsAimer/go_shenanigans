package main
import "fmt"

func main(){
	var a,b = return_nums()
	fmt.Println(a,b)

	//_ is bland identifier and is used when a value isnt used
	c,_ := return_nums()
	_=c

}

func return_nums () (int, int){
	return 1,2
}