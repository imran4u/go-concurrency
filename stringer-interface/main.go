package main

import "fmt"

type user struct {
	name  string
	email string
}

// enbuilt stringer interface fuction String() string
func (u user) String() string {
	return fmt.Sprintf("name:%s, email:%s", u.name, u.email)
}

func main() {
	u := user{
		name:  "imran",
		email: "imran@eamil.com",
	}

	fmt.Println(u) // before override stringer interface out put : {imran imran@eamil.com}
	// after override out put : name:imran, email:imran@eamil.com
}
