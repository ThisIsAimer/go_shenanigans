package main
import "fmt"

const discription = "participant details are: "
const(
	monday = 1
	tuesday int = 2
)

var(
	thursday = 3
	wednesday = 4
)

func main(){
	var age int = 20
	var age1 = 19
	var name = "andro"
	name1 := "potato"
	var boolean = true


	fmt.Println(discription+name,age,",",name1,age1,!boolean)
	printgame()

}

func printgame(){
	var game = "chess"
	fmt.Println(game, monday,thursday)
}