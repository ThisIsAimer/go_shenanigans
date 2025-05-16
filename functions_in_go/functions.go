package main
import "fmt"

func main(){
	var a,b = return_nums()
	fmt.Println(a,b)

}

func return_nums () (int, int){
	return 1,2
}