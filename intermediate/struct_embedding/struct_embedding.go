package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	employeeInfo person // Embedded struct Named field
	// person  // Anonymous field
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.empId, e.salary)
}

func main() {

	emp := Employee{
		employeeInfo: person{name: "John", age: 30},
		empId:        "E001", salary: 50000,
	}

	fmt.Println("Name:", emp.employeeInfo.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", emp.employeeInfo.age)   // Same as above
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	emp.introduce()

}