package main
import "fmt"

func main(){
	decimel := 10
	fmt.Printf("|%10d|\n",decimel)
	fmt.Printf("|%-10d|\n",decimel)

	var text = "hello"

	fmt.Printf("|%10s|\n",text)
	fmt.Printf("|%-10s|\n",text)



}