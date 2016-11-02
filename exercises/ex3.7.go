/*
Modify the code to use a pointer but still be able to initialize without using the dot notation.
*/

package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func main() {
	p := Player{&User{55, "Joe", "Melrose"}, 22}
	fmt.Println(p.Greetings())
}
