package server

//CircularQueue is a circular queue
type CircularQueue struct {
	Queue     []interface{}
	HeadIndex int
	len       int
}

func newCircularQueue(len int) *CircularQueue {
	return &CircularQueue{
		Queue:     make([]interface{}, len),
		HeadIndex: -1,
		len:       len,
	}
}

func (c *CircularQueue) push(x interface{}) {
	c.HeadIndex = (c.HeadIndex + 1) % c.len
	c.Queue[c.HeadIndex] = x
}

func (c *CircularQueue) getAll() []interface{} {
	return c.Queue
}
