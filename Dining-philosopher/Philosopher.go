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

	//Continous loop where philosophers try to receive values from forks in order to deem if they are in use.
	//If no value can be received from the forks channels, the default cases are executed, and the loop continues.

	for {
		//trying to receive a value from channel.
		//If channel is ready left fork is avaliable If not do nothing(jumps to default case line 52)
		select {
		case <-p.fork_left.inUse:
			select {
			//Try to receive vale from right fork and if channel is not ready jump to default case and release leftfork(line 48)
			case <-p.fork_right.inUse:
				//if both forks are free, eat and release forks afterwards.
				p.timesEating++
				fmt.Printf("%s is currently eating and has eaten: %d time(s)\n", p.Name, p.timesEating)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(4500)))
				// release fork by sending a value to the forks channel
				p.fork_right.inUse <- false
				fmt.Printf("%s is thinking \n", p.Name)

			default:
			}
			//release a fork by sending a value to the forks channel
			p.fork_left.inUse <- false
		default:
		}
	}
}
