package circularfifoqueue

/*
CircularFifoQueue is a first-in first-out queue with a fixed size that replaces its oldest element if full.
The reference for this package was taken from
https://commons.apache.org/proper/commons-collections/apidocs/src-html/org/apache/commons/collections4/queue/CircularFifoQueue.html
*/

type CircularFifoQueue struct {
	head   int
	tail   int
	buffer []interface{}
	full   bool
}

func NewCircularFifoQueue(size int) *CircularFifoQueue {
	buffer := make([]interface{}, size)
	start, end := 0, 0
	return &CircularFifoQueue{head: start, tail: end, buffer: buffer}
}

/*
Returns the maxium capacity of the queue
*/
func (q *CircularFifoQueue) Capacity() int {
	return len(q.buffer)
}

/*
Returns the number of elements stored in the queue
*/
func (q *CircularFifoQueue) Len() int {
	if q.tail < q.head {
		return q.Capacity() - q.head + q.tail
	} else if q.head == q.tail {
		if q.full {
			return q.Capacity()
		}
		return 0
	} else {
		return q.tail - q.head
	}
}

/*
Enqueue a value into the queue. If the queue is full
it will replace the oldest element.
*/
func (q *CircularFifoQueue) Enqueue(i interface{}) {
	isFull := q.Capacity() == q.Len()
	if isFull {
		q.Dequeue()
	}
	q.buffer[q.tail] = i
	q.tail = (q.tail + 1) % q.Capacity()
	if q.head == q.tail {
		q.full = true
	}
}

/*
Dequeue a value from the Ring buffer.
Returns nil if the ring buffer is empty.
*/
func (q *CircularFifoQueue) Dequeue() interface{} {
	v := q.buffer[q.head]
	if v != nil {
		q.buffer[q.head] = nil
		q.head = (q.head + 1) % q.Capacity()
		q.full = false
	}
	return v
}

/*
Rather than try and implement an iterator that might not suite
every single need we make available the underlying buffer.
Warning: Do not modify anything here and use it for read only purposes.
The following link has a good rundown of different iterator patterns
http://ewencp.org/blog/golang-iterators/
*/
func (q *CircularFifoQueue) Values() []interface{} {
	return q.buffer
}

/*
Returns the initial index in the buffer of the queue
*/
func (q *CircularFifoQueue) Head() int {
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
	return q.tail
}
