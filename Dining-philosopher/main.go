package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"asger", "john", "Annna", "bjørn", "ogge"}

	philospherArr := make([]philosopher, 0, 5)

	forkArr := make([]fork, 0, 5)

	for i := 0; i < 5; i++ {
		philosopher := makePhilosopher(arr[i])

		philospherArr = append(philospherArr, *philosopher)

		forks := create_fork(i)

		forkArr = append(forkArr, *forks)

	}
	fmt.Println(philospherArr)

	for i := 0; i < len(forkArr); i++ {
		go grant_fork(forkArr[i])
	}

	fmt.Println(forkArr)

}
