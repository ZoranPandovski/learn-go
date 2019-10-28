# Concurrency in Go

### Week 1

Parallel Execution

- two programs executing at exactly the same time
- needs replicated hardware
- tasks may complete more quickly

Concurrent Execution

- concurrent execution is not necessarily the same as parallel execution
- Concurrent - start and end times overlap
- Parallel - execute at exactly the same time
- may or may not be executed on the same hardware
- programmer determines which tasks can be executed in parallel
- mapping tasks to hardware - OS and Go runtime scheduler
- concurrency improves performance even without parallelism - hiding latency
- other concurrent tasks can operate while one task is waiting

### Week 2

Processes - an instance of a running program

Things unique to a process:

1. Memory - virtual address space, code, stack, heap, shared libraries
2. Registers - program counter, data registers, stack pointer

OS - allows many processes to execute concurrently
Processes are switched quickly - 20ms
User has the impression of parallelism.
OS must give processes fair access to resources.
A process can have multiple threads.
Threads share some context.
OS schedules threads rather than processes nowdays.
Goroutines are like a thread in Go.
Multiple Goroutines execute within a single OS thread.

Go Runtime Scheduler

- schedules goroutines inside an OS thread
- like a little OS inside a single OS thread
- logical processor is mapped to a thread

Interleavings

- order of execution within a task is known
- order of execution between concurrent tasks is unknown

Race conditions occur due to communication. If a variable is shared between their contexts.

### Week 3

One goroutine is created automatically to execute the main().
Other goroutines are created using the 'go' keyword.

    // main goroutine blocks on call to foo()
    a = 1
    foo()
    a = 2
    
    // new goroutine created for foo(). main does not block
    a = 1
    go foo()
    a = 2

goroutines exit when its code is complete.
all goroutines exit when main goroutine is complete. (early exit)
Delayed Exit Hack - time.Sleep(100 * time.Millisecond)

Synchronization - using global events whose execution is viewed by all threads simultaneously

Sync Wait Group

- sync package contains functions to synchronize between goroutines
- sync.WaitGroup forces a goroutine to wait for oter goroutines
- contains an internal counter to keep track

    func foo(wg *sync.WaitGroup) {
    	fmt.Println("New routine")
    	wg.Done()
    }
    
    func main() {
    	var wg sync.WaitGroup
    	wg.Add(1)
    	go foo(&wg)
    	wg.Wait()
    	fmt.Print("Main routine")
    }
    
    /*
    	Add() - increments the counter
    	Wait() - blocks until counter == 0
    	Done() - decrements the counter
    */

Goroutine communication

- goroutines usually work together to perform a bigger task
- often need to send data to collaborate

Channels

- transfer data between goroutines
- channels are typed
- use make() to create a channel, `c := make(chan int)`
- send and receive data using `<-` operator
- send data on a channel `c <- 3`
- receive data from a channel `a := <- c`

    func prod(v1 int, v2 int, c chan int) {
    	c <- v1 * v2
    }
    func main() {
    	c := make(chan int)
    	go prod(1, 2, c)
    	go prod(3, 4, c)
    	a := <- c
    	b := <- c
    	fmt.Print(a * b)
    }

Default channels are unbuffered.
Unbuffered channels cannot hold data in transit.
Sending blocks until data is received.
Receiving blocks until data is sent.

Blocking & Synchronization

1. Channel communication is synchronous.
2. Blocking is the same as waiting for communication.
3. Receiving and ignoring the result is same as a Wait().

Channel Capacity

- channels can contain a limited number of objects
- default size if 0, i.e. unbuffered
- capacity is the number of objects it can hold in transit
- optional arguments to make defines a channel capacity `c := make(chan int, 3)`
- sending only blocks if buffer is full
- receiving only blocks if buffer is empty
- use of buffering - sender and receiver do not need to operate at exactly the same speed

### Week 4

Iterating through a channel:

- common to iteratively read from a channel for `i := range c { fmt.Println(i) }`
- continues to read from channel `c`
- one iteration each time a new value is received
- `i` is assigned to the read value
- iterates when sender calls `close(c)`

Receiving from multiple goroutines:

    a := <- c1
    b := <- c2
    fmt.Println(a * b)

Select Statement - choose FCFS data from a set of channels

    select {-
    	case a = <- c1:
    		fmt.Print(a)
    	case b = <- c2:
    		fmt.Print(b)
    }

Select Send or Receive - may select either send or receive operations

    select {
    	case a = <- inchan:
    		fmt.Print("Received a")
    	case outchan <- b:
    		fmt.Print("Sent b")
    }

Select with an Abort channel:

- use select with a separate abort channel
- may want to receive data until an abort signal is received

    for {
    	select {
    		case a = <- c:
    			fmt.Print(a)
    		case <- abort:
    			return
    	}
    }

Default Select - may want a default operation to avoid blocking

    select {-
    	case a = <- c1:
    		fmt.Print(a)
    	case b = <- c2:
    		fmt.Print(b)
    	default:
    		fmt.Print("nop")
    }

Concurrency is at the machine code level.

Sync.Mutex - a mutex ensures mutual exclusion. Uses a binary semaphore.

Lock() - puts the shared variable flag up.
If a lock is already taken by a goroutine, Lock() blocks until the flag is put down.

Unlock() - puts the shared variable flag down.
When called a bocked Lock() can be processed.

    var i int = 0
    var mut sync.Mutex
    func inc() {
    	mut.Lock()
    	i = i + 1
    	mut.Unlock()
    }

Synchronous Initialization

Sync.Once has one method `once.Do(f)`.
Function `f` is executed only one time even if tit is called in multiple goroutines.
All calls to once.Do() block until the first returns. Ensures that initialization executes first.

    var on sync.Once
    func setup() {
    	fmt.Println("Init")
    }
    func dostuff() {
    	on.Do(setup)
    	fmt.Println("hello")
    	wg.Done()
    }
