package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Concurrency in go: Do not communicate by sharing memory; instead, share memory by communicating.
 */
func main() {
	// goroutine is a lightweight thread managed by Go runtime
	// execution of say in following statement occurs in separate goroutine
	go say("Go")
	// following function is executed in current goroutine i.e. main goroutine
	say("Hello")

	// Declaration of channel
	ch := make(chan int)

	// Use slice instead of array in GOlang
	a := []int{12, 23, 78, -9, 12, -34, 7, 89}

	// channel is a stack so works on FIFO queue principle
	// Note that goroutine is not pushed on stack
	// goroutines are intiated sequentially don't confuse with deferred functions
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	// Note the assignment operator is := before channel operator <-
	first, second := <-ch, <-ch
	fmt.Println(first, second)

	// buffered channels can be used to make other goroutines wait for the current goroutine to get
	// finished with its work
	// ** Note that their declaration is similar to slices with len & cap but only with length as
	// it doesnot have underlying array
	chBuffered := make(chan int, 2)
	// sender1 is main goroutine
	chBuffered <- 12
	chBuffered <- 23
	// Deadlock condition 1: Pushing value beyond channel capacity within the same goroutine
	// Following statement will result in the deadlock as we are pushing the values beyond the capacity
	// in the main goroutine(leighweight thread)
	// chBuffered <- 45
	// use of closure & go invocation immidiately will result in concurrent goroutine
	// Note: it will have access to channel without passing channel variable explicitly
	go func() {
		// sender2 is this subroutine
		chBuffered <- 45
		close(chBuffered)
		//Following statement won't be seen on console because of low sync?
		// fmt.Println(<-chBuffered)
	}()
	// for v := range chBuffered {
	// 	fmt.Println(v)
	// }
	// make the current goroutine to add wait time for other goroutine
	// time.Sleep(100 * time.Millisecond)
	fmt.Println(<-chBuffered)
	fmt.Println(<-chBuffered)
	fmt.Println(<-chBuffered)
	// Deadlock condition 2: Retrive value from non closed empty channel
	// Following statement will result in the deadlock if not closed by sender
	// if closed variable will be initialised to its zero value
	v, ok := <-chBuffered
	if ok {
		fmt.Println(v)
	}

	// Buffered channel has ony capacity no length
	ch1 := make(chan int, 10)
	fmt.Println("Length = ", len(ch1), "Capacity = ", cap(ch1))
	go fibonacci(cap(ch1), ch1)
	fmt.Println("Length = ", len(ch1), "Capacity = ", cap(ch1))
	// channel range like array range
	for i := range ch1 {
		fmt.Println(i)
	}

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)

	//Creation of ticking & boom channels
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(600 * time.Millisecond)

	for i := 0; i < 10; i++ {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			// on "boom" return from main function
			// if no return used will continue all the for loop iterations
			// return
		default:
			fmt.Println(" .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	genc := gen(2, 3, 4)
	out := sq(genc)
	// Last stage of pipiline i.e. consumer
	// If we refer only one channel out of above two channels will get the proper output
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	// Compose the functions sq & gen
	// Multiple compositions possible for functions with same type of inboud and outbound channel
	for i := range sq(sq(gen(6, 2, 11))) {
		fmt.Println(i)
	}

	// Fan out: Distribution of outbound channel over different functions
	genc1 := gen(3, 4)
	t1 := sq(genc1)
	t2 := sq(genc1)
	// fmt.Println(t1, t2)

	// The merge function converts a list of channels to a single channel by starting a goroutine
	// for each inbound channel that copies the values to the sole outbound channel. Once all the output
	// goroutines have been started, merge starts one more goroutine to close the outbound channel after
	// all sends on that channel are done.
	for n := range merge(t1, t2) {
		fmt.Println(n)
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	// This is a closure which takes an argument so defined as separate function first
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func say(message string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}

// Channels are a typed conduit i.e. it carry a specific type of data e.g. int, float etc.
// channel operator "<-" is used to assign value into channel or recieve value from channel
func sum(a []int, s chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	s <- sum
	// Following statement will print for the first go routine only due to very minute sync time
	// This is why time package is used to add the delay before such operation
	fmt.Println(sum)
}

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		// Concurrent assignment not sequential assignment
		a, b = b, a+b
	}
	// Need to close it here else range in for loop will exceed the capacity of the channel
	close(ch)
}

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		// Select makes the goroutine with for loop wait as that goroutine makes use of the
		// channel vaules to be set by the select block
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// **Pipelines**
// They most useful mechanisms for maintaining a sequence in asynchronous goroutines
// Stage: It is a group of goroutines running the same function
// First Stage(Producer): It has only outbound channels
func gen(list ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// second stage: It has inbound as well as outbound channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// for loop range argument won't pull the pipe contents out
		// instead traverse the channel like an array UNTIL it is CLOSED
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
