package main

import "fmt"


/* 
	A race condition occurs when two or more threads can access shared data and they try to change it at the same time. 
	Because the thread scheduling algorithm can swap between threads at any time,
	you don't know the order in which the threads will attempt to access the shared data.
	Therefore, the result of the change in data is dependent on the thread scheduling algorithm,
	i.e. both threads are "racing" to access/change the data.
*/
func main() {
	a := 0
	for i := 0; i < 100; i++ {
		go func() {
			a++
		}()
	}
	y := a + 5
	fmt.Printf("a: %d\n", y)
}

/*
Race Condition : A race condition or race hazard is the condition of an electronics,
software, or other system where the system's substantive behavior is dependent
on the sequence or timing of other uncontrollable events.
It becomes a bug when one or more of the possible behaviors is undesirable.
Why Race condition occurs?
It occurs when two or more process can access and change the shared data at the same time.
It occurred because there were conflicting accesses to a resource .
Critical section problem may cause race condition.
*/