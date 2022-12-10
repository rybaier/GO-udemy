package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // dont have to declare an field if embedding struct
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// ways to declare a struct
	// alex := person{"Alex", "Anderson"}
	// fmt.Println(alex)
	// joe := person{firstName: "Joe", lastName: "Smith"}
	// fmt.Println(joe)
	// var larry person
	// larry.firstName = "Larry"
	// larry.lastName = "unknown"
	// fmt.Println(larry)
	// fmt.Printf("%+v", larry) //prints both keys and values
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	// // fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()

	// mySlice := []string{"hi", "there", "how", "are", "you"}
	// updateSlice(mySlice)
	// fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
func (p person) print() {
	fmt.Printf("%+v", p)
}
