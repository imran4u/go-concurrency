package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t) //2024-12-28 07:57:05.050435977 +0530 IST m=+0.000012015
	//format
	fmt.Println(t.Format("02-01-2006"))                     //28-12-2024
	fmt.Println(t.Format("Monday 02-01-2006"))              //Saturday 28-12-2024
	fmt.Println(t.Format("Monday 02-Jan-2006 03-04-05 PM")) //Saturday 28-Dec-2024 07-58-55 AM
	// Notice 15 here for 24 hour format
	fmt.Println(t.Format("Monday 02-Jan-2006 15-04-05 ")) //Saturday 28-Dec-2024 07-59-44

	// Reverse parsting.
	st := "Saturday 28-12-2024"

	fmt.Println(time.Parse("Monday 02-01-2006", st))

}
