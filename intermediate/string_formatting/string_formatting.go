package main
import "fmt"

func main(){
	decimel := 10
	fmt.Printf("|%10d|\n",decimel)
	fmt.Printf("|%-10d|\n",decimel)

	var text = "hello"

	fmt.Printf("|%10s|\n",text)
	fmt.Printf("|%-10s|\n",text)

	number := 20.5343535
	fmt.Printf("number is %.4f\n", number)

	var tag = `"number is %.4f\n", number`
	fmt.Println(tag) //whole thing taken as string


}