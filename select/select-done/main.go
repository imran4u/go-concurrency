package main

import (
	"fmt"
	"net/http"
)

func isExist(done <-chan struct{}, urls ...string) (results chan *http.Response, errs chan error) {
	result := make(chan *http.Response)
	err := make(chan error)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				return
			default:
				respose, e := http.Get(url)
				if e != nil {
					err <- e
					continue
				}
				result <- respose
			}
		}

		close(result) // close the chanel
		close(err)// close the chanel

	}()

	return result, err
}

func main() {
	done := make(chan struct{})
	res, err := isExist(done, "http://google.com", "http://localhost")

	for range 2 {
		select {
		case r := <-res:
			fmt.Println(r.StatusCode)
		case e := <-err:
			fmt.Println(e)
		}
	}
	close(done) // close the chanel after use

}
