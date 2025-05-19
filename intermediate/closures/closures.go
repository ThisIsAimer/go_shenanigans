package main
import "fmt"


// cloasures is when a referenced function can access the variables of
//the outer functions. which helps in encapsulation.

func main(){

	var sequence = adder()

	println("value of i is now",sequence())
	println("value of i is now",sequence())
	println("value of i is now",sequence())

	var sequence2 = adder()
	println("value of i is now",sequence2())
	println("1 value of i is now",sequence())


	var subtractor = func() func(int) int{
		var i = 100
		return func(j int) int {
			return i - j
		}
	}()

	println("subtracted value is: ", subtractor(10))
	println("subtracted value is: ", subtractor(20))
	

}

func adder() func() int{
	var i = 0
	fmt.Println("value of i is", i)
	return func() int{
		i++
		return i
	}
}