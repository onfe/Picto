package server

//CircularQueue is a circular queue
type CircularQueue struct {
	queue     []interface{} `json:"Queue"`
	HeadIndex int           `json:"HeadIndex"`
	len       int
}

func newCircularQueue(len int) *CircularQueue {
	return &CircularQueue{
		queue:     make([]interface{}, len),
		HeadIndex: -1,
		len:       len,
	}
}

func (c *CircularQueue) push(x interface{}) {
	c.HeadIndex = (c.HeadIndex + 1) % c.len
	c.queue[c.HeadIndex] = x
}

func (c *CircularQueue) getAll() []interface{} {
	return c.queue
}
