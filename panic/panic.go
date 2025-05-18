package main
import "fmt"

func main(){
	positive_inputs(-10)

}


//as no code can be passed after panic, all defer will be passed before it occurs
func positive_inputs(number int) {
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")

	if number < 0 {
		panic("negetive number")
	}
	defer fmt.Println("third defer")

	fmt.Println("the number is:", number)
}