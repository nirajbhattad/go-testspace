# Scheduler

In Go, goroutines are lightweight threads that are managed by the Go runtime. The Go runtime uses a scheduler to manage the execution of goroutines. The scheduler is responsible for allocating CPU time to different goroutines, and it uses a number of algorithms to determine which goroutines should run and when.


- The scheduler uses a **work-stealing algorithm**, which allows goroutines to steal work from other goroutines that have completed their current task. 
- The scheduler also uses a **priority-based algorithm**, which allows goroutines with higher priority to be scheduled before goroutines with lower priority. 
- Additionally, the scheduler uses a **load-balancing algorithm**, which distributes the workload among available CPU cores.

The scheduler also uses a technique called **"Goroutine preemption"** which allows the scheduler to interrupt the execution of a goroutine and schedule another one if it's deemed necessary. This allows the scheduler to prevent a single goroutine from monopolizing the CPU resources.


Here is an example of how the scheduler works for goroutines in Go:

```go

package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("goroutine 1:", i)
            time.Sleep(time.Millisecond * 100)
        }
    }()
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("goroutine 2:", i)
            time.Sleep(time.Millisecond * 100)
        }
    }()
    time.Sleep(time.Second * 2)
}

```

In this example, two goroutines are created and run concurrently. Each goroutine has a loop that prints out a message and sleeps for 100 milliseconds. The main goroutine also sleeps for 2 seconds, allowing the other goroutines to run.

When the program is executed, the scheduler will allocate **CPU time** to both goroutines, allowing them to run concurrently. The exact order in which the messages are printed will depend on the **scheduling algorithm** used by the scheduler.

In summary, the Go runtime scheduler is responsible for allocating CPU time to different goroutines. The scheduler uses a number of algorithms, such as **work-stealing, priority-based, and load-balancing algorithms,** to determine which goroutines should run and when. The scheduler also uses a technique called "Goroutine preemption" to interrupt the execution of a goroutine and schedule another one if necessary, this allows the scheduler to prevent a single goroutine from monopolizing the CPU resources.

## Scheduler Runtime

Go will handle multiplexing goroutines onto OS threads for you.

The algorithm it uses to do this is known as a work `stealing strategy`.

**fair scheduling/ load balancing** :  In an effort to ensure all processors were equally utilized, we could evenly distribute the load between all available processors. Imagine there are n processors and x tasks to perform. In the fair scheduling strategy, each processor would get x/n tasks:

Go models concurrency using a fork-join model.

As a refresher, remember that Go follows a fork-join model for concurrency. Forks are when goroutines are started, and join points are when two or more goroutines are synchronized through channels or types in the sync package. The work stealing algorithm follows a few basic rules. Given a thread of execution:

At a fork point, add tasks to the tail of the deque associated with the thread.

**Go scheduler’s job is to distribute runnable goroutines over multiple worker OS threads that runs on one or more processors.** In multi-threaded computation, two paradigms have emerged in scheduling: work sharing and work stealing.

- Work-sharing: When a processor generates new threads, it attempts to migrate some of them to the other processors with the hopes of them being utilized by the idle/underutilized processors.
- Work-stealing: An underutilized processor actively looks for other processor’s threads and “steal” some.

The migration of threads occurs less frequently with work stealing than with work sharing. When all processors have work to run, no threads are being migrated. And as soon as there is an idle processor, migration is considered.

**Go has a work-stealing scheduler since 1.1**, contributed by Dmitry Vyukov. This article will go in depth explaining what work-stealing schedulers are and how Go implements one.


**Scheduling basics**

Go has an M:N scheduler that can also utilize multiple processors. At any time, M goroutines need to be scheduled on N OS threads that runs on at most GOMAXPROCS numbers of processors. Go scheduler uses the following terminology for goroutines, threads and processors:

- G: goroutine<br/>
- M: OS thread (machine)<br/>
- P: processor<br/>

There is a P-specific local and a global goroutine queue. Each M should be assigned to a P. Ps may have no Ms if they are blocked or in a system call. At any time, there are at most GOMAXPROCS number of P. At any time, only one M can run per P. More Ms can be created by the scheduler if required.

**Why have a scheduler?**

goroutines are user-space threads
conceptually similar to kernel threads managed by the OS, but managed entirely by the Go runtime

* lighter-weight  and cheaper than kernel threads.

* smaller memory footprint:
    * initial goroutine stack = 2KB; default thread stack = 8KB
    * state tracking overhead
    * faster creation, destruction, context switchesL
    * goroutines switches = ~tens of ns; thread switches = ~ a us.

