package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://invalid-url",
	}
	for _, url := range urls {
		url := url // capture variable make local copy on each iteration
		g.Go(func() error {
			req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
			_, err := http.DefaultClient.Do(req)
			if err != nil {
				return fmt.Errorf("failed: %s", url)
			}
			fmt.Println("success:", url)
			return nil
		})
	}

	// Note:  if any of go-routine failed and return error all other will canceled.

	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
	}

}
