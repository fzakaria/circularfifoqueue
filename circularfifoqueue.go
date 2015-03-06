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
Peek at the value at the head. It will return Nil
if the queue is empty
*/
func (q *CircularFifoQueue) Peek(i interface{}) interface{} {
	return q.buffer[q.head]
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
Iterate through the queue in FIFO order via callback.
The choice of callback was chosen based on the evaluation done
in the following link http://ewencp.org/blog/golang-iterators/
*/
func (q *CircularFifoQueue) Do(f func(interface{})) {
	isFirst := q.full
	index := q.head
	for isFirst || index != q.tail {
		f(q.buffer[index])
		index = (index + 1) % q.Capacity()
		isFirst = false
	}
}

/*
Returns a copy of the underlying buffer.
You are free to modify this as necessary
*/
func (q *CircularFifoQueue) Values() []interface{} {
	b := make([]interface{}, q.Capacity())
	copy(b, q.buffer)
	return b
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
