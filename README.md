# go-micro-patterns

A collection of frequently used Go patterns for micro services and concurrency.


## ForkJoinGroup

A facade over sync.WaitGroup to encapsulate the error-prone `wg.Add(1)`, `go func() {defer wg.Done()}` pattern.

