package queue

type Q interface {
	Push(data interface{})
	Pop() interface{}
	Len() int
}

func NewQueue() Q {
	return &queue{}
}

type queue struct {
	q  []interface{}
	size int
}

func (q *queue) Push(data interface{}) {
	q.q = append(q.q, data)
	q.size++
}

func (q *queue) Pop() interface{} {
	if q.size <= 0 {
		return nil
	}
	tmp := q.q[0]
	q.q = q.q[1:]
	q.size--
	return tmp
}

func (q *queue) Len() int {
	return q.size
}


