package server

//CircularQueue is a circular queue
type CircularQueue struct {
	queue     []interface{}
	headIndex int
	len       int
}

func newCircularQueue(len int) CircularQueue {
	return CircularQueue{
		queue:     make([]interface{}, len),
		headIndex: 0,
		len:       len,
	}
}

func (c *CircularQueue) push(x interface{}) {
	c.queue[c.headIndex] = x
	c.headIndex = (c.headIndex + 1) % c.len
}

func (c *CircularQueue) getAll() []interface{} {
	return c.queue
}
