package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {

	dinner := Dinner{[]*philosopher{}, []*fork{}, [5]string{}}

	Make_dinner(&dinner)

	wg.Add(15)
	dinner.Start()
	wg.Wait()

}
