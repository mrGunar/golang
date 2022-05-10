package main

import "fmt"

func (u User) getAllInfo() {
	fmt.Println(u)
	for _, x := range u.hobbies {
		fmt.Println(x)
	}

	switch {
	case u.status == true:
		fmt.Printf("Status: %s", u.last_name)
	case u.status == false:
		fmt.Printf("Status: %s", u.name)
	}
}

type User struct {
	name      string
	last_name string
	age       uint16
	city      string
	male      string
	status    bool
	hobbies   []string
}

func main() {

	user := User{"Alex", "Smith", 20, "Brighton", "m", true, []string{"footbal", "cookies", "guitar"}}
	user.getAllInfo()
}
