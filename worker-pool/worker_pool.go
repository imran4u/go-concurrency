package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Id int
}

type Result struct {
	Id    int
	Value int // just double the id for demo
}

func Worker(workerId int, job chan Job, result chan Result, w *sync.WaitGroup) {
	defer w.Done()
	for j := range job {
		r := Result{
			Id:    j.Id,
			Value: j.Id * 2, // just mock
		}
		result <- r
	}

}

func main() {
	jobCount := 5
	workerCount := 3

	jobs := make(chan Job, jobCount)
	results := make(chan Result, jobCount)
	wg := sync.WaitGroup{}
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go Worker(i, jobs, results, &wg)
	}

	for i := 1; i <= jobCount; i++ {
		jobs <- Job{Id: i}
	}
	close(jobs)
	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result : ", r.Id, ":", r.Value)
	}

}
