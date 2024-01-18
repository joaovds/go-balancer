package balancer

import "sync"

type Balancer struct {
  workers []*Worker
  JobQueue chan int
  Wg *sync.WaitGroup
}

func NewBalancer(numberOfWorkers int) *Balancer {
  JobQueue := make(chan int)
  workers := make([]*Worker, numberOfWorkers)
  Wg := sync.WaitGroup{}

  for i := 0; i < numberOfWorkers; i++ {
    worker := NewWorker(i, JobQueue, &Wg)
    workers[i] = worker
  }
  Wg.Add(numberOfWorkers)

  return &Balancer{workers, JobQueue, &Wg}
}

func (b *Balancer) Start(verbose bool) {
  for _, worker := range b.workers {
    go worker.Start(verbose)
  }
}
