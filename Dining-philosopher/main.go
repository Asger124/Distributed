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

// 	arr := [5]string{"asger", "john", "Annna", "bjørn", "ogge"}

// 	philospherArr := make([]*philosopher, 0, 5)

// 	forkArr := make([]*fork, 0, 5)

// 	for i := 0; i < 5; i++ {
// 		forks := create_fork(i)

// 		forkArr = append(forkArr, forks)

// 	}
// 	fmt.Println("first philo :", philospherArr)

// 	for i := 0; i < len(forkArr); i++ {

// 		philosopher := makePhilosopher(arr[i], forkArr[i], (forkArr[(i+1)%5]))

// 		philospherArr = append(philospherArr, philosopher)

// 		//philospherArr[i].fork_right.Occupied = true
// 	}
// 	wg.Add(10)
// 	for i := 0; i < 5; i++ {

// 		go func(i int) {
// 			defer wg.Done()
// 			forkArr[i].grant_fork()
// 		}(i)

// 		go func(i int) {
// 			defer wg.Done()
// 			philospherArr[i].InitiatePhilosopher()
// 		}(i)
// 	}

// 	wg.Wait()

// 	fmt.Println("first Forkarr : ", forkArr)
// 	fmt.Println("second philoarr :", philospherArr)
// 	fmt.Println("second forkarr :", forkArr)

// 	for i := range forkArr {
// 		fmt.Println(*forkArr[i])
// 		fmt.Println("filo forks:", philospherArr[i].fork_right)
// 	}
// }
