package balancer

import (
	"log"
	"sync"
	"time"
)

type Worker struct {
  id int
  jobQueue chan int
  Wg *sync.WaitGroup
}

func NewWorker(id int, jobQueue chan int, wg *sync.WaitGroup) *Worker {
  return &Worker{id: id, jobQueue: jobQueue, Wg: wg}
}

func (w *Worker) Start(verbose bool) {
  for job := range w.jobQueue {
    defer time.Sleep(time.Second)
    if verbose {
      log.Println("worker", w.id, "finished job", job + 1)
    }
  }

  w.Wg.Done()
}

