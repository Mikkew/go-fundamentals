package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Profile struct {
	Bio string
}

type User struct {
	Name    string
	Addr    Address
	Profile *Profile //Opcional
}

func main() {
	user := User{
		Name: "Ana",
		Addr: Address{
			City:  "Buenos Aires",
			State: "Ciudad Autónoma de Buenos Aires",
		},
	}

	if user.Profile == nil {
		fmt.Println("Sin perfil")
	}

	user.Profile = &Profile{
		Bio: "Soy Ana y me gusta Go",
	}

	fmt.Println("Bio: ", user.Profile.Bio)
	fmt.Println(user)
}
