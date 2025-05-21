package intermediate

import "fmt"

func main() {

	//Printing Functions
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Formatting Functions
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	// Scanning Functions
	var name1 string
	var age1 int

	fmt.Print("Enter your name and age:")
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	fmt.Scanf("%s %d", &name1, &age1)
	fmt.Printf("Name: %s, Age: %d\n", name1, age1)

	// Error Formatting Functions

	err := checkAge(19)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}