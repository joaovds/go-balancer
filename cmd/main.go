package main

import (
	"github.com/joaovds/go-balancer/internal/balancer"
	"github.com/joaovds/go-balancer/internal/flags"
)

func main() {
  instanceFlags := flags.SetupFlags()

  workersQuantity := 1
  balance := balancer.NewBalancer(workersQuantity)

  balance.Start(*instanceFlags.Verbose)

  for i := 0; i < 15; i++ {
    balance.JobQueue <- i
  }

  close(balance.JobQueue)
  balance.Wg.Wait()
}
