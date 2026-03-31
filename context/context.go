package main

import "fmt"

func main() {

	type myKey string
	type asKey = string // as type
	var sKey string

	mkey := myKey("key")
	aKey := asKey("key")
	key := "key"
	sKey = "key"

	if sKey == key {
		fmt.Println("sKey == key ", sKey == key)
	}

	// can't use mkey == key , custom type != string, you must have to convert one, use this concept to pass conext key in context.
	if mkey == myKey(key) {
		fmt.Println("mkey == myKey(key) || string(mkey) == key ", string(mkey) == key)
	}

	//No need to case in case of as key
	if aKey == key {
		fmt.Println("aKey == key ", aKey == key)
	}

}
