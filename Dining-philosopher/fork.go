package main

type fork struct {
	id       int
	inUse    chan bool
	Occupied bool
}

func create_fork(id int) *fork {
	f := fork{}
	f.id = id
	f.inUse = make(chan bool)
	f.Occupied = false

	return &f

}

// Continous loop which attempts send and receive boolean value through channel f.inUse.
// Firstly the function will attempt to send a value through the channel and the loop halts until another go routine is ready to receive
// the value from the channel. When another go routine has received the value, the loop will halt once again unntil another go routine,
// sends a value into the variable f.Occuped.
// One iteration in the loop entails that a philosopher has picked up a fork, and then put it down again
// Essentially if a fork running a in GO routine, is picked up, it is blocked until it is put down again.

func (f *fork) grant_fork() {
	for {
		//Attempts to send a boolean value through the channel and blocks until another go routine receives the value.
		f.inUse <- f.Occupied
		// Attempts to receive a value and blocks until it receives a value from another go routine.
		f.Occupied = <-f.inUse
	}

}
