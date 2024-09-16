# Dining Philospherrs

# HOW TO RUN THE PROGRAM 

# From the terminal navigate to the folder \Distributed\Dining-philosopher. 
# When in the folder Dining-philosopher, write: go run . in the terminal. 
# the program will run indefinetly - to exit the program press Ctrl + C. 


# HOW DEADLOCK IS AVOIDED 

# Deadlock is avoided by using the waitgroups. The main go routine is added, and the waitgroup
# ensures that the go routine never stops running.

# Deadlock is also avoided by the implementation of InitiatePhilosopher() in Philosopher.go(Line 27) and the implementatin of grant_fork(), in fork.go (Line 26). The default cases in the select statements in the InititatePhilosopher() method helps to ensure that the philosophers do not pick up a fork and hold onto it while waiting for another. if they cant get a hold of both forks they will release any forks,
# they hold and try again later - due to the default cases. The use of unbuffered channels and the infinite for loop in grant_fork()
# will ensure that if a fork is picked up it is 'blocked' until it is put down again. Which helps to ensure proper syncronization.
# The methods are explained further in their respective files.
