# circularfifoqueue

[![GoDoc](https://godoc.org/github.com/fzakaria/circularfifoqueue?status.svg)](https://godoc.org/github.com/fzakaria/circularfifoqueue)

--
    import "github.com/fzakaria/circularfifoqueue"

CircularFifoQueue is a first-in first-out queue with a fixed size that replaces its oldest element if full.

The package is intentially kept very minimal. The use case of this package as opposed to `ring` defined by Go is that it provides access to the underlying slice. This is a BYOI (Bring Your Own Iterator) solution and you can implement many solutions as a result with the slice (too difficult to solve all needs).




