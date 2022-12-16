package main

import "fmt"

func main(){
	n := 100
	jobs := make(chan int, n)
	results := make(chan int, n)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	

	for i := 0; i < n; i++ {
		jobs <- i
	}
	close(jobs)



	for j := 0; j < n; j++ {
		fmt.Println(<- results)
	}
}

func worker(jobs <- chan int, results chan<- int){
	for n := range jobs{
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}