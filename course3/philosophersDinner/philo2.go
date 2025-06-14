package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Implement the dining philosopher’s problem with the following constraints/modifications.
// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// The host allows no more than 2 philosophers to eat concurrently.
// Each philosopher is numbered, 1 through 5.
// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

type Chopstick struct {
	mut sync.Mutex
}

type Philosopher struct {
	number int
	leftChopstick *Chopstick
	rightChopstick *Chopstick
}

type Host struct {
    permission chan struct{} // Use struct{} as a semaphore
}


func (p Philosopher) eat(wg *sync.WaitGroup, host Host) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        // Request permission (blocks if 2 are eating)
        host.permission <- struct{}{}

        // Randomize chopstick pickup order to prevent deadlock
        if rand.Intn(2) == 0 {
            p.leftChopstick.mut.Lock()
            p.rightChopstick.mut.Lock()
        } else {
            p.rightChopstick.mut.Lock()
            p.leftChopstick.mut.Lock()
        }

        fmt.Println("starting to eat", p.number)
        fmt.Println("finishing eating", p.number)

        p.rightChopstick.mut.Unlock()
        p.leftChopstick.mut.Unlock()

        // Release permission
        <-host.permission
    }
}

func hostGoroutine(wg *sync.WaitGroup) {
    defer wg.Done()
    // Host just manages the channel, which inherently limits to 2 due to buffer size
    // No additional logic needed here since the channel acts as the host's control
    wg.Wait() // Wait for all philosophers to finish
}

func main() {
    // Initialize chopsticks
    chopsticks := make([]*Chopstick, 5)
    for i := 0; i < 5; i++ {
        chopsticks[i] = new(Chopstick)
    }

    // Initialize philosophers
    philosophers := make([]*Philosopher, 5)
    for i := 0; i < 5; i++ {
        philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
    }

    // Initialize host with capacity 2
    host := Host{make(chan struct{}, 2)}

    // WaitGroup for synchronization
    var wg sync.WaitGroup
    wg.Add(5) // 5 philosophers

    // Launch host goroutine
    var hostWg sync.WaitGroup
    hostWg.Add(1)
    go hostGoroutine(&hostWg)

    // Launch philosopher goroutines
    for i := 0; i < 5; i++ {
        go philosophers[i].eat(&wg, host)
    }

    // Wait for all philosophers to finish
    wg.Wait()
    hostWg.Wait()
}