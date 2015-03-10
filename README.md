# circularfifoqueue

[![GoDoc](https://godoc.org/github.com/fzakaria/circularfifoqueue?status.svg)](https://godoc.org/github.com/fzakaria/circularfifoqueue)

--
    import "github.com/fzakaria/circularfifoqueue"

CircularFifoQueue is a first-in first-out queue with a fixed size that replaces its oldest element if full. Access to the queue is goroutine safe and is protected via RWMutex.

The package is intentially kept very minimal. The use case of this package as opposed to `ring` defined by Go is that it provides access to the underlying slice and implements the Queue interface.




