// An "embedded" struct is just a struct that actas like a field inside another struct.(nested)
package main

import "fmt"

func main() {
	type Contatc struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name       string
		salary     int
		contatInfo Contatc
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contatInfo: Contatc{
			email:   "jkeller@at.com",
			address: "Streen 20, London",
			phone:   00230023230,
		},
	}
	fmt.Printf("%+v\n", john)

	fmt.Printf("%v\n", john.contatInfo.email)

	john.contatInfo.email = "jkLL@cc.com"
	fmt.Printf("%v\n", john.contatInfo.email)

	myContact := Contatc{email: "joe@cc.com", phone: 2343421, address: "Grace, 23 Ourinhs"}
	_ = myContact
}
