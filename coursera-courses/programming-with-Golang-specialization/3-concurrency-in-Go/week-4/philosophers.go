package main

import (
	"fmt"
	"sync"
	"time"
)
/**********
	** Implement the dining philosopher’s problem with the following constraints/modifications.

		1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
		2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
		3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
		4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		5. The host allows no more than 2 philosophers to eat concurrently.
		6. Each philosopher is numbered, 1 through 5.
		7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, 
			where <number> is the number of the philosopher.
		8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
			where <number> is the number of the philosopher.
***********/

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id, EatenTimes int
	leftChop, rightChop *Chopstick
}

func (phil Philosopher) Eat(wg *sync.WaitGroup, c chan *Philosopher) {

	for i := 0; i < 3; i++ {
		c <- &phil

		if phil.EatenTimes < 3 {
			phil.leftChop.Lock()
			phil.rightChop.Lock()
			fmt.Printf("Starting to eat %d\n", phil.id)

			phil.EatenTimes++

			fmt.Printf("Finishing eating %d\n", phil.id)
			phil.rightChop.Unlock()
			phil.leftChop.Unlock()
			wg.Done()
		}
	}
}

func host(wg *sync.WaitGroup, c chan *Philosopher) {
	for {
		if len(c) == 2 {
			<- c
			<- c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	// init wait group
	var wg sync.WaitGroup

	// init 5 Chopsticks 
	chopsticks := make([]*Chopstick, 5)

	// init 5 Philosophers
	philosophers :=  make([]*Philosopher, 5)

	// init a channel
	c := make(chan *Philosopher, 2)

	wg.Add(15) // 5 Philosophers with 3 EatenTimes

	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{}
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i + 1, // id
			0, // EatenTimes
			chopsticks[i], // left chopstick
			chopsticks[(i + 1) % 5], // right chopstick
		}
	}

	go host(&wg, c)

	for _, phil := range philosophers {
		go phil.Eat(&wg, c)
	}

	wg.Wait()
}
