package main

import (
	"context"
	"fmt"
	"net/http"
)

type Result struct {
	code int
	e    error
}

// func checkUrl(ctx context.Context, url string, result chan Result) {
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Early exit url : ", url)
// 		return
// 	default:
// 		res, err := http.Get(url)
// 		if err != nil {
// 			result <- Result{
// 				code: 0,
// 				e:    err,
// 			}
// 		} else {
// 			result <- Result{
// 				code: res.StatusCode,
// 			}
// 		}

// 	}

// }

func checkUrl(ctx context.Context, url string, result chan Result) {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- Result{e: err}
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		select {
		case result <- Result{e: err}:
		case <-ctx.Done():
		}
		return
	}

	defer res.Body.Close()

	select {
	case result <- Result{code: res.StatusCode}:
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan Result)

	urls := []string{
		"http://localhost",
		"http://Google.com",
		"https://github.com",
		"https://golang.org",
	}

	for _, url := range urls {
		go checkUrl(ctx, url, result)
	}

	for i := 0; i < len(urls); i++ {
		r := <-result
		if r.e != nil {
			cancel() // cancel all the go routine
		}
		fmt.Println(r.code, ":", r.e)
	}

}
