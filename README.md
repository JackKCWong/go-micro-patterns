# go-micro-patterns

A collection of frequently used Go patterns for micro services and concurrency.

`go get github.com/JackKCWong/go-micro-patterns`

## `concur` concurrency patterns
### ForkJoinGroup

```golang
import "github.com/JackKCWong/go-micro-patterns/concur"
```

A facade over sync.WaitGroup to encapsulate the error-prone `wg.Add(1)`, `go func() {defer wg.Done()}` pattern.

## `cnp` cloud native patterns
### Retry / RetryWithInterval / RetryWithBackoff

```golang
import "github.com/JackKCWong/go-micro-patterns/cnp"
```

