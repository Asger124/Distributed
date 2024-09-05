package main

type fork struct {
	id       int
	isfree   chan bool
	Occupied bool
}

func create_fork(id int) *fork {
	f := fork{}
	f.id = id
	f.isfree = make(chan bool)
	f.Occupied = false

	return &f

}

func grant_fork(f fork) {

	f.isfree <- f.Occupied
	f.Occupied = <-f.isfree
}

// 	if !f.isfree {
// 		fmt.Printf("fork request failed. fork is occupied by %s\n", f.occupied_by)
// 		return p.fork_left

// 	} else {
// 		f.isfree = false
// 		f.occupied_by = p.Name
// 	}
// 	return f
// }
