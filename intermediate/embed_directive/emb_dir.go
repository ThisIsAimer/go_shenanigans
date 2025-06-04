package main
import (
	"fmt"
	_ "embed"
)


//must be in the same folder

//go:embed example.txt
var content string

func main(){
	fmt.Println("embeded content:",content)

}