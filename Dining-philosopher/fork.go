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

func (f *fork) grant_fork() {
	for {
		f.inUse <- f.Occupied
		f.Occupied = <-f.inUse
	}

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
