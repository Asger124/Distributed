package main

type Dinner struct {
	Philosophers []*philosopher
	Forks        []*fork
	Names        [5]string
}

func Make_dinner(dinner *Dinner) {
	dinner.Names = [5]string{"Asger", "Jovana", "Jojo", "Marco", "Thore"}

	for i := 0; i < len(dinner.Names); i++ {
		forks_for_dinner := create_fork(i)

		dinner.Forks = append(dinner.Forks, forks_for_dinner)

	}

	for i := 0; i < len(dinner.Names); i++ {

		philosopher := makePhilosopher(dinner.Names[i], dinner.Forks[i], (dinner.Forks[(i+1)%(len(dinner.Names))]))

		dinner.Philosophers = append(dinner.Philosophers, philosopher)
	}

}

func (dinner *Dinner) Start() {

	for i := 0; i < len(dinner.Names); i++ {

		go dinner.Forks[i].grant_fork()

	}

	for i := 0; i < len(dinner.Names); i++ {

		go dinner.Philosophers[i].InitiatePhilosopher()

	}

}
