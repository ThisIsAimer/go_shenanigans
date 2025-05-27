package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main(){
	programme := 1

	if programme == 0{
	//rand.intn generates a random number by itself from 0 to the given number
	//autoseeding
	var randomNumber = rand.Intn(101)

	fmt.Println("random number:", randomNumber)
	//--------------------------------------------------------------------------
	//using seed
	//NewSource creates a source with a deterministic seed (7)
	//rand.New(...) uses that source to create a new random number generator
	var seed = rand.New(rand.NewSource(7))

	fmt.Println("random number with seed:",seed.Intn(101))

	//fmt.Println("float rand:", math.Round(rand.Float64()*10000)/100)
	fmt.Println("float rand:",rand.Float64()*100)// between 0.0 to 1.0
	} else{
		var choice int
			for{
			fmt.Println("hey welcome to the dice roll")
			fmt.Println("press 1 to roll dice")
			fmt.Println("press 2 to exit")
			fmt.Print("enter: ")

			var _,err = fmt.Scanln(&choice)
			if err != nil{
				fmt.Println(err)
				return
			}
			if choice == 1{
				break
			}else if choice == 2{
				fmt.Print("closing....")
				return
			}else{
				fmt.Println("please enter one of the numbers!")
			}
		}
		if choice == 1{
			for {
				var dice1 = rand.Intn(6)+1
				var dice2 = rand.Intn(6)+1
				fmt.Println("------------------------------------------------")
				fmt.Printf("Number in the first dice is %d\nNumber in the second dice is %d\n", dice1,dice2)
				fmt.Printf("The total you rolled is: %d\n", dice1+dice2)
				fmt.Println("------------------------------------------------")
				fmt.Println("do you want to roll again Y/N")
				var choice1 string
				var _,err = fmt.Scanln(&choice1)
				if err != nil{
					fmt.Println(err)
					return
				}
				choice1 = strings.ToLower(choice1)
				if choice1 == "y"{
					continue
				} else if choice1 == "n"{
					fmt.Print("closing....")
					break
				}else{
					for{
						fmt.Println("please only type y or n:")
						var _,err = fmt.Scanln(&choice1)
						if err != nil{
							fmt.Println(err)
							return
						}
						if choice1 == "y"{
							break
						} else if choice1 =="n"{
							fmt.Print("closing....")
							return
						}else{}
					}
				}
			}
		}
	}

}
