package main
import "fmt"

type Rectangle struct{
	length float64
	width float64
}

func (r Rectangle) New(l float64, w float64) Rectangle{
	return Rectangle{length: l, width: w}
}

func main(){
	var rectangle = Rectangle{}.New(10,20)
	fmt.Println(rectangle)

}