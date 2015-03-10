package circularfifoqueue

import "sync"

/*
CircularFifoQueue is a first-in first-out queue with a fixed size that replaces its oldest element if full.
The reference for this package was taken from
https://commons.apache.org/proper/commons-collections/apidocs/src-html/org/apache/commons/collections4/queue/CircularFifoQueue.html
*/

type CircularFifoQueue struct {
	sync.RWMutex
	head   int //head index of the first element in the queue
	tail   int //tail - 1 is the index of the last element in the queue
	buffer []interface{}
	full   bool //flag set when we have just appended and filled the queue (distinguish between empty)
}

func NewCircularFifoQueue(size int) *CircularFifoQueue {
	if size <= 0 {
		panic("CircularFifoQueue: Size must be greater or equal to 1")
	}
	buffer := make([]interface{}, size)
	start, end := 0, 0
	return &CircularFifoQueue{head: start, tail: end, buffer: buffer}
}

/*
Returns the maxium capacity of the queue
*/
func (q *CircularFifoQueue) Capacity() int {
	q.RLock()
	defer q.RUnlock()
	return q.capacity()
}

/*
Helper function to be called by others that doesn't lock
*/
func (q *CircularFifoQueue) capacity() int {
	return len(q.buffer)
}

/*
Returns the number of elements stored in the queue
*/
func (q *CircularFifoQueue) Len() int {
	q.RLock()
	defer q.RUnlock()
	return q.len()
}

/*
Helper function to be called by others that doesn't lock
*/
func (q *CircularFifoQueue) len() int {
	if q.tail < q.head {
		return q.capacity() - q.head + q.tail
	} else if q.head == q.tail {
		if q.full {
			return q.capacity()
		}
		return 0
	} else {
		return q.tail - q.head
	}
}

/*
Peek at the value at the head. It will return Nil
if the queue is empty
*/
func (q *CircularFifoQueue) Peek(i interface{}) interface{} {
	q.RLock()
	defer q.RUnlock()
	return q.buffer[q.head]
}

/*
Enqueue a value into the queue. If the queue is full
it will replace the oldest element.
*/
func (q *CircularFifoQueue) Enqueue(i interface{}) {
	q.Lock()
	defer q.Unlock()

	isFull := q.capacity() == q.len()
	if isFull {
		q.dequeue()
	}
	q.buffer[q.tail] = i
	q.tail = (q.tail + 1) % q.capacity()
	if q.head == q.tail {
		q.full = true
	}
}

/*
Dequeue a value from the Ring buffer.
Returns nil if the ring buffer is empty.
*/
func (q *CircularFifoQueue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()
	return q.dequeue()
}

/*
Helper function to be called by others that doesn't lock
*/
func (q *CircularFifoQueue) dequeue() interface{} {
	v := q.buffer[q.head]
	if v != nil {
		q.buffer[q.head] = nil
		q.head = (q.head + 1) % q.capacity()
		q.full = false
	}
	return v
}

/*
Iterate through the queue in FIFO order via callback.
start: the initial index (modulo capacity) to start iterating
func(interface{}) bool: the callback that will receive each item.
												Returning False stops iterations.
The choice of callback was chosen based on the evaluation done
in the following link http://ewencp.org/blog/golang-iterators/
*/
func (q *CircularFifoQueue) Do(start int, f func(interface{}) bool) {
	q.RLock()
	defer q.RUnlock()

	start = start % q.capacity()

	isFirst := q.full
	index := q.head
	for isFirst || index != q.tail {
		if !f(q.buffer[index]) {
			break
		}
		index = (index + 1) % q.capacity()
		isFirst = false
	}
}

/*
Returns a copy of the underlying buffer.
You are free to modify this as necessary
*/
func (q *CircularFifoQueue) Values() []interface{} {
	q.RLock()
	defer q.RUnlock()
	b := make([]interface{}, q.capacity())
	copy(b, q.buffer)
	return b
}

/*
Returns the initial index in the buffer of the queue
*/
func (q *CircularFifoQueue) Head() int {
	q.RLock()
	defer q.RUnlock()
	return q.head
}

/*
Returns the last index in the bufer of the queue
Index mod Len() of the array position following the last queue
element.  Queue elements start at elements[start] and "wrap around"
elements[maxElements-1], ending at elements[decrement(end)].
For example, elements = {c,a,b}, start=1, end=1 corresponds to
the queue [a,b,c].
*/
func (q *CircularFifoQueue) Tail() int {
	q.RLock()
	defer q.RUnlock()
	return q.tail
}
