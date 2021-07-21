package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	Number int
	sync.Mutex
}

type Philosopher struct {
	Number         int
	LeftChopstick  *ChopStick
	RightChopstick *ChopStick
	EatenTimes     int
}

/* 
	this not my implementations I just keep it here for the sake of learning
*/

// Eat method tries to eat 3 times by asking host if it can. When it's done it finishes with a Done() call to the wait group
func (p *Philosopher) Eat(wg *sync.WaitGroup, canEat func() bool, doneEat func()) {
	for p.EatenTimes < 3 {
		if !canEat() {
			continue
		}
		p.LeftChopstick.Mutex.Lock()
		p.RightChopstick.Mutex.Lock()

		fmt.Printf("Starting to eat %d\n", p.Number)
		p.EatenTimes += 1

		fmt.Printf("Finishing eating %d\n", p.Number)

		p.LeftChopstick.Mutex.Unlock()
		p.RightChopstick.Mutex.Unlock()
		doneEat()
	}
	wg.Done()
}

func main() {

	// Create sticks
	var sticks []*ChopStick
	for i := 0; i < 5; i++ {
		chop := ChopStick{
			Number: i + 1,
		}
		sticks = append(sticks, &chop)
	}

	// Create Philosophers
	var philos []*Philosopher
	for i := 0; i < 5; i++ {
		philo := Philosopher{
			Number:         i + 1,
			LeftChopstick:  sticks[i],
			RightChopstick: sticks[(i+1)%5],
			EatenTimes:     0,
		}
		philos = append(philos, &philo)
	}

	// Mutex for host to control the number of Philosophers currently eating. Ensures only 2 eat at a time
	var mut sync.Mutex
	var currentlyEating int

	// canEatFunc checks if there are less than 2 philosophers currently busy eating. If so allow another to eat, if not dissallow
	canEatFunc := func() bool {
		mut.Lock()
		defer mut.Unlock()
		if currentlyEating < 2 {
			currentlyEating += 1
			return true
		} else {
			return false
		}
	}

	// Frees up a new Philosopher to eat
	doneEatFunc := func() {
		mut.Lock()
		currentlyEating -= 1
		mut.Unlock()
	}

	// Start goroutines
	var wg sync.WaitGroup
	for _, p := range philos {
		wg.Add(1)
		go p.Eat(&wg, canEatFunc, doneEatFunc)
	}

	// Wait until all Philosophers have eaten 3 times
	wg.Wait()
}