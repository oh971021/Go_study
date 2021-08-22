package main

import (
	"fmt"
	"sync"
	"time"
)

type job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	// 10개의 작업을 할당한다.
	fmt.Printf("%d 작업 시작\n", j.index)

	time.Sleep(1 * time.Second)

	// 10개의 작업 결과를 출력한다.
	fmt.Printf("%d 작업 완료 - 결과 %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
