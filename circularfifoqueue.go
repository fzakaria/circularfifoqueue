package circularfifoqueue

/*
CircularFifoQueue is a first-in first-out queue with a fixed size that replaces its oldest element if full.
*/

type CircularFifoQueue struct {
	start  int
	end    int
	buffer []interface{}
}

func NewCircularFifoQueue(size int) *CircularFifoQueue {
	buffer := make([]interface{}, size)
	start, end := 0, 0
	return &CircularFifoQueue{start: start, end: end, buffer: buffer}
}

/*
Returns the maxium capacity of the queue
*/
func (q *CircularFifoQueue) Len() int {
	return len(q.buffer)
}

/*
Enqueue a value into the queue. If the queue is full
it will replace the oldest element.
*/
func (q *CircularFifoQueue) Enqueue(i interface{}) {
	q.buffer[q.end] = i
	q.end = (q.end + 1) % q.Len()
	//if we've looped aroudn then also increase the head
	if q.head == q.end {
		q.head = (q.head + 1) % q.Len()
	}
}

/*
Dequeue a value from the Ring buffer.
Returns nil if the ring buffer is empty.
*/
func (q *CircularFifoQueue) Dequeue() interface{} {
	v := q.buffer[q.head]
	q.head = (q.head + 1) % q.Len()
	//if we've looped aroudn then also increase the tail
	if q.head == q.end {
		q.tail = (q.tail + 1) % q.Len()
	}
	return v
}

/*
Rather than try and implement an iterator that might not suite
every single need we make available the underlying buffer.
Warning: Do not modify anything here and use it for read only purposes.
The following link has a good rundown of different iterator patterns
@link http://ewencp.org/blog/golang-iterators/
*/
func (q *CircularFifoQueue) Values() []interface{} {
	return q.buffer
}

/*
Returns the initial index in the buffer of the queue
*/
func (q *CircularFifoQueue) Start() int {
	return q.start
}

/*
Returns the last index in the bufer of the queue
*/
func (q *CircularFifoQueue) End() int {
	return q.end
}
