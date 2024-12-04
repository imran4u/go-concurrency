package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

var person Person
var once sync.Once
var wg sync.WaitGroup

func instancePerson(name string, age int) {
	once.Do(func() {
		person = Person{name: name, age: age}
	},
	)

}

func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			instancePerson(fmt.Sprintf("ali %d ", i), i) // person ali -> a (exchange)
		}()
	}

	//instancePerson("ali %d ", 27) // person ali -> a (exchange)
	// instancePerson("farhan", 25) // person farhan -> b eidgah
	wg.Wait()
	fmt.Println(person)

}
