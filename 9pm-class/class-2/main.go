package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	time1 := time.Now().UnixMilli()
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.whatsapp.com",
		"https://www.twitter.com",
		"https://www.amazon.com",
		"https://www.flipkart.com",
	}

	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		go downloadSite(url, &wg)
	}
	wg.Wait()
	fmt.Println("Time taken ", time.Now().UnixMilli()-time1)
}

func downloadSite(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "Fail error: ", err)

		return
	}

	fmt.Println(url, "success")

	bs := make([]byte, 9999)
	res.Body.Read(bs)
	fmt.Println(string(bs))

}
