package main

import (
	"context"
	"fmt"
)

// Avoid key collision in context value.
func main() {
	type mykey string

	ctx := context.Background()

	cv := context.WithValue(ctx, "userId", 123) // parent

	cv = context.WithValue(cv, "userId", 149) // new , Note key collition

	cv2 := context.WithValue(cv, mykey("userId"), 189) // no key collition as it is diffent type.
	value := cv.Value("userId")
	fmt.Println("value=", value)

	// cv2
	cv2Value1 := cv2.Value("userId")
	cv2Value2 := cv2.Value(mykey("userId"))
	fmt.Println("cv2Value1 =", cv2Value1)
	fmt.Println("cv2Value2 =", cv2Value2)

}
