package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

var wg sync.WaitGroup

type ChopS struct {
	mut sync.Mutex
	id int
}

type Philo struct {
	leftCS, rightCS *ChopS
	id int
}

func lockMutexes(c1 *ChopS, c2 *ChopS) {
	if rand.Intn(2)%2 == 0 {
		c1.mut.Lock()
		c2.mut.Lock()
	} else {
		c2.mut.Lock()
		c1.mut.Lock()
	}
}

func unlockMutexes(c1 *ChopS, c2 *ChopS) {
	if rand.Intn(2)%2 == 0 {
		c2.mut.Unlock()
		c1.mut.Unlock()
	} else {
		c1.mut.Unlock()
		c2.mut.Unlock()
	}
}

var eating = 0
var mutEating sync.Mutex

func (p Philo) hostAllowToEat() bool {
	mutEating.Lock()
	if eating < 2 {
		mutEating.Unlock()
		return true
	} else {
		mutEating.Unlock()
		return false
	}
}


func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		if (p.hostAllowToEat()) {
			mutEating.Lock()
			eating++
			mutEating.Unlock()

			lockMutexes(p.leftCS, p.rightCS)

			fmt.Printf("starting to eat %d\n", p.id)
			// Eating...
			time.Sleep(2 * time.Second)
			fmt.Printf("finishing eating %d\n", p.id)

			unlockMutexes(p.leftCS, p.rightCS)
			mutEating.Lock()
			eating--
			mutEating.Unlock()
		} else {
			i--
		}
	}
	fmt.Printf("\n")
	wg.Done()
}


func main() {
	var n int = 5
	CSticks := make([]*ChopS, n)
	for i:= 0; i < n; i++ {
		CSticks[i] = &ChopS{
			id: i,
		}
	}

	philos := make([]*Philo, n)
	for i:= 0; i < n; i++ {
		philos[i] = &Philo{
			leftCS: CSticks[i],
			rightCS: CSticks[(i + 1)%n],
			id: i,
		}
	}

	for i:= 0; i<n; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
