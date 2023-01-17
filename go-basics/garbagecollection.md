# Garbage Collection


In Go, garbage collection is the process of automatically freeing memory that is no longer needed by the program.
The Go runtime uses a **concurrent garbage collector** that runs in parallel with the main program. 
This means that the garbage collector can reclaim memory while the program is still running,
which minimizes the impact of collection pauses on the performance of the program.

The garbage collector uses a technique called **"tracing"** to determine which objects are still reachable and which are not.
The garbage collector starts by marking the objects that are directly reachable from the program's roots (e.g. global variables, stack frames, etc.).
Then, it follows the references from these objects to find other reachable objects, and so on, until it has marked all the reachable objects. 
The objects that are not marked are considered unreachable and are eligible for collection.

The Go runtime uses a combination of techniques such as **mark-sweep, mark-compact, and incremental collection** to efficiently manage memory.

The Go garbage collector uses a concurrent, tri-color, mark-and-sweep algorithm. The algorithm is concurrent,
meaning that it runs in the background while the program is running, and it uses a tri-color marking scheme to keep
track of memory that is in use and memory that is no longer needed. The algorithm uses a "stop-the-world" phase to stop program execution 
and mark the memory that is in use. Once the marking phase is complete, the algorithm then begins the sweep phase, during which it frees 
memory that is no longer in use.

In Go, the garbage collector uses a technique called **"generational garbage collection"** to improve performance. 
The heap is split into two regions, one for new objects and another for long-lived objects, called **generations.** 
The garbage collector will collect the new object region more frequently than the long-lived region.
This is because the new object region is more likely to have more garbage than the long-lived region.

In summary, Go uses a concurrent, tri-color, **mark-and-sweep garbage collector**, which periodically 
checks for memory that is no longer needed by the program and frees it. Additionally, the garbage collector uses a technique 
called **"generational garbage collection" **

Here's an example that illustrates the use of garbage collection in Go:

```go

package main

func main() {
    var m runtime.MemStats
    var a [1024]byte
    runtime.ReadMemStats(&m)
    fmt.Println("Allocated memory before:", m.Alloc)
    for i := 0; i < 100000000; i++ {
        a = [1024]byte{}
    }
    runtime.ReadMemStats(&m)
    fmt.Println("Allocated memory after:", m.Alloc)
}

```

- In this example, the program creates an array a of 1024 bytes and assigns it to a variable. 
The program then creates the array 100000000 times, but the garbage collector will reclaim the memory allocated for the previous array after 
it is no longer needed. You can see the allocated memory before and after the loop,
showing how the garbage collector is able to reclaim the memory that is no longer needed.

- It's important to notice that the garbage collector runs automatically and you do not need to explicitly invoke it.
However, you can control the garbage collection with the GOGC environment variable, which sets the initial garbage collection target percentage. 
Also, you can use the debug.SetGCPercent function to change the garbage collection target percentage programmatically.

