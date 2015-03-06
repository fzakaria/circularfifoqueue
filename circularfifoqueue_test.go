package circularfifoqueue

import "testing"

func TestBasicRemoveNoInsert(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 5; i++ {
		x := q.Dequeue()
		if x != nil {
			t.Fatal("Unexpected response", x, "wanted", nil)
		}
	}
}

func TestBasicInsertNoRollover(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		x := q.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestBasicInsertRemoveNoRollover(t *testing.T) {
	q := NewCircularFifoQueue(20)
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		q.Dequeue()
	}
	for i := 5; i < 10; i++ {
		x := q.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

//[ 10 , 11 , 12 , 13 , 14 , 5 , 6, 7 , 8, 9]
func TestBasicInsertRollover(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 15; i++ {
		q.Enqueue(i)
	}
	for i := 5; i < 15; i++ {
		x := q.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

//[ 10 , 11 , 12 , 13 , 14 , 5 , 6, 7 , 8, 9]
//[10 , 11,  12, 13, 14, nil, nil, nil, 8, 9 ]
func TestBasicInsertRemoveRollover(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 15; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		q.Dequeue()
	}
	for i := 8; i < 15; i++ {
		x := q.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestBasicRemoveTooMuch(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		x := q.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
	for i := 0; i < 5; i++ {
		x := q.Dequeue()
		if x != nil {
			t.Fatal("Unexpected response", x, "wanted", nil)
		}
	}
	if q.Head() != q.Tail() {
		t.Fatal("Unexpected response", q.Head(), "wanted", q.Tail())
	}
}

func TestBasicLenNoRollover(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	if q.Len() != 5 {
		t.Fatal("Unexpected response", q.Len(), "wanted", 5)
	}
}

func TestBasicLenEmpty(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		q.Dequeue()
	}
	if q.Len() != 0 {
		t.Fatal("Unexpected response", q.Len(), "wanted", 0)
	}
}

//[10 , 11,  12, 13, 14, nil, nil, nil, 8, 9 ]
func TestBasicLenRollover(t *testing.T) {
	q := NewCircularFifoQueue(10)
	for i := 0; i < 15; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		q.Dequeue()
	}
	if q.Len() != 7 {
		t.Fatal("Unexpected response", q.Len(), "wanted", 7)
	}
}
