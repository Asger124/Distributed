package main

type philosopher struct {
	Name        string
	fork_left   fork
	fork_right  fork
	timesEating int
	state       string
}

func makePhilosopher(name string) *philosopher {

	p := philosopher{}
	p.Name = name
	p.timesEating = 0
	return &p

}
