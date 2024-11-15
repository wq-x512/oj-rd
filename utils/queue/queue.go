package queue

type Queue struct {
	items chan string
}

func NewQueue(capacity int) *Queue {
	Q := &Queue{items: make(chan string, capacity)}
	Q.Enqueue("http://222.186.61.53:8000")
	Q.Enqueue("http://150.158.107.119:8000")
	Q.Enqueue("http://222.186.61.53:8001")
	Q.Enqueue("http://150.158.107.119:8001")
	Q.Enqueue("http://222.186.61.53:8002")
	Q.Enqueue("http://150.158.107.119:8002")
	return Q
}
func (q *Queue) Enqueue(value string) {
	q.items <- value
}
func (q *Queue) GetOne() string {
	value := <-q.items
	q.items <- value
	return value
}
func (q *Queue) Dequeue() string {
	return <-q.items
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
