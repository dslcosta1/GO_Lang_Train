package main

import (
	"fmt"
	"sync"
)

/* Because there are 5 sticks and each philosopher needs 2
sticks for eating it is impossible to have more than 2
philosophers eating at the same time, so the condition to limit
the permissions from the host to 2 is useless, although it
will be implemented */
/* we will use 2 channels with capacity 2 to communicate
philosophers with thehost. There will be 2 tokens in the
Request channel where philosophers can get permission from the host
When each philosopher finishes eating send the token back in the Answer
channel.
The Host function read the token in the Answer channel
and put it back in the Request channel
We have selected an integer as a Token */

var eatWg sync.WaitGroup

// type chopstick
type ChopS struct{ sync.Mutex }

// type Philosopher
type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

// function eat
func (p Philo) eat(c1 chan int, c2 chan int) {
	var token int // to read token fron channel1
	defer eatWg.Done()
	for j := 0; j < 3; j++ {

		if p.id == 0 {
			p.rightCS.Lock()
			p.leftCS.Lock()
		} else {
			p.leftCS.Lock()
			p.rightCS.Lock()
		}
		token = <-c1 //read free token from channel 1
		fmt.Printf("Philosopher #%d is eating iteration %d\n", p.id, j)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Printf("Philosopher #%d has finished eating iteration %d\n", p.id, j)
		c2 <- token // free token to channel 2
	}
}

// Host read tokens from the answer channel when a philosopher
// has finished eating and then it adds it to the request channel
// to be available for other philosophers
func Host(c1 chan int, c2 chan int) {
	var token int
	for {
		token = <-c2
		c1 <- token
	}
}

func main() {

	// Initializaton of chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	// Initializaton of philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	// use channels for giving permission to 2 philosophers
	Requestc := make(chan int, 2)
	Answerc := make(chan int, 2)

	// initialization of Requestc with 2 tokens of permissions
	Requestc <- 1
	Requestc <- 1
	//start host function
	go Host(Requestc, Answerc)

	// start eating
	for i := 0; i < 5; i++ {
		eatWg.Add(1)
		go philos[i].eat(Requestc, Answerc)
	}
	eatWg.Wait() // wait until all philo have eaten
}