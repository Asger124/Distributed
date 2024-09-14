package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	dinner := Dinner{[]*philosopher{}, []*fork{}, [5]string{}}

	Make_dinner(&dinner)

	wg.Add(1)
	dinner.Start()
	fmt.Println("philosphers", dinner.Philosophers[0].fork_left.Occupied, dinner.Philosophers[0].fork_right.Occupied)
	wg.Wait()

}
