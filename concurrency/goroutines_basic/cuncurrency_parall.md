- **CONCURRENCY** means loading more goroutines at a time. If one goroutine blocks, another one is picked up and started. On single core CPU you can run ONLY concurrent applications but they are not run in parallel.
- **PARALLELISM** means multiple goroutines execute at the same time. It requires multiple CPUs.
- Concurrency means independently executing processes or dealing with multiple things at once, while parallelism is the simultaneous execution of processes and require multiple core CPUs.

# Goroutines
- A *goroutine* is a lightweight thread of execution;
- A goroutine is a function that is capable of running concurrently with other functions.
- Goroutine are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes a fixed size of 1-2Mb.
- Goroutine stack size shrinks and grows as needed.
- **Scheduling a goroutine is much cheaper than scheduling a thread.**
- Os threads are scheduled by the OS kernel (slow operation due to the number of mem access required).
    But goroutines are scheduled by its own GO Scheduler using a technique called **m:n scheduling**


