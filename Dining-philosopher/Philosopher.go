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

	//Infinite loop where philosophers try to receive values from forks in order to deem if they are free.
	//If no value can be received from the forks channels, the default cases are executed, and the loop continues.

	for {
		//trying to receive a value from channel.
		//If channel is ready left fork is avaliable If not do nothing(jumps to default case line 51)
		select {
		case <-p.fork_left.inUse:
			select {
			//Try to receive vale from right fork and if channel is not ready jump to default case and release leftfork(line 47)
			case <-p.fork_right.inUse:
				//if both forks are free eat and, release forks afterwards.
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

// for i := 0; p.timesEating < 3; i++ {

// 	//Time delay to ensure that the philosophers engage in dining at different random times.

// 	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))

// 	// If both forks are free, engage in eating

// 	if !<-p.fork_left.inUse && !<-p.fork_right.inUse {
// 		fmt.Printf("%s have arrived", p.Name)

// 		p.timesEating++
// 		fmt.Printf("%s Are currently eating and has eaten %d times\n", p.Name, p.timesEating)

// 		fmt.Println("forks when eating", p.fork_right.Occupied, p.fork_left.Occupied)

// 		//Simulate time a philosopher has eating
// 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(4500)))

// 		if p.timesEating == 3 {
// 			fmt.Printf("%s Is now full\n", p.Name)
// 		}

// 		//Drop both forks after eating

// 		p.fork_left.inUse <- false
// 		p.fork_right.inUse <- false

// 		fmt.Printf("%s Has started thinking\n", p.Name)

// 		fmt.Println("left forks after :", p.fork_left.Occupied, p.fork_right.Occupied)

// 	}
// 	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

// 	wg.Done()

// }
