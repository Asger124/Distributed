package main

import (
	"fmt"
	"math/rand"
	"time"
)

type philosopher struct {
	Name        string
	fork_left   *fork
	fork_right  *fork
	timesEating int
}

func makePhilosopher(name string, left, right *fork) *philosopher {

	p := philosopher{}
	p.Name = name
	p.timesEating = 0
	p.fork_left = left
	p.fork_right = right
	return &p

}

func (p *philosopher) InitiatePhilosopher() {

	for i := 0; p.timesEating < 3; i++ {

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))

		if !<-p.fork_left.inUse && !<-p.fork_right.inUse {

			//p.fork_left.inUse <- true

			//stateright := <-p.fork_right.inUse

			// if !<-p.fork_right.inUse {

			// 	p.fork_right.inUse <- true

			p.timesEating++
			fmt.Printf("%s Are currently eating and has eaten %d times\n", p.Name, p.timesEating)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(4500)))

			if p.timesEating == 3 {
				fmt.Printf("%s Is now full\n", p.Name)
			}

			p.fork_right.inUse <- false
			p.fork_left.inUse <- false
			fmt.Printf("%s Has started thinking\n", p.Name)

		} else {
			if <-p.fork_left.inUse {
				p.fork_left.inUse <- false
			}
			if <-p.fork_right.inUse {
				p.fork_right.inUse <- false
			}

			// Wait for a short time before retrying
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}

		wg.Done()

	}
}
