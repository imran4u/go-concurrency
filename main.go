// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 2, 5)
// 	fmt.Println("len : ", len(slice), "Cap : ", cap(slice))

// 	slice1 := new([]int)
// 	fmt.Println("len : ", len(*slice1), "Cap : ", cap(*slice1))

// }

package main

import "fmt"

// Base struct (can be embedded in other structs)
type Engine struct {
	HorsePower int
}

func (e *Engine) Start() {
	fmt.Println("Engine starting...")
}

func (e *Engine) Stop() {
	fmt.Println("Engine stopping...")
}

// Car struct embeds Engine struct
type Car struct {
	Engine // Embedded struct
	Model  string
}

func (c *Car) Drive() {
	fmt.Println("Driving", c.Model)
}

func main() {

	car := Car{
		Engine: Engine{HorsePower: 200}, // Embed Engine
		Model:  "Tesla Model S",
	}

	car.Start() // Calls Start() from the embedded Engine struct
	car.Drive() // Calls Drive() from Car struct
	car.Stop()  // Calls Stop() from the embedded Engine struct
}
