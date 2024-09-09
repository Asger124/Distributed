package main

type Dinner struct {
	P     []*philosopher
	F     []*fork
	Names [5]string
}

func Make_dinner(d *Dinner) {
	d.Names = [5]string{"Asger", "Jovana", "Bjærn", "Anna", "Ogge"}

	for i := 0; i < len(d.Names); i++ {
		forks := create_fork(i)

		d.F = append(d.F, forks)

	}

	for i := 0; i < len(d.Names); i++ {

		philosopher := makePhilosopher(d.Names[i], d.F[i], (d.F[(i+1)%(len(d.Names))]))

		d.P = append(d.P, philosopher)
	}

}

func (d *Dinner) Start() {

	for i := 0; i < len(d.Names); i++ {

		go d.F[i].grant_fork()

	}

	for i := 0; i < len(d.Names); i++ {

		go d.P[i].InitiatePhilosopher()

	}

}
