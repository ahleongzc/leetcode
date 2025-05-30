package stack

// This implementation's pop operation takes up O(n) while the push operation takes up O(1)
// You can swap the pop operation to take up O(1) while the push operation takes up O(n) as well

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	// Enqueue operation
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	for _ = range len(this.queue) - 1 {
		// Dequeue operation
		el := this.queue[0]
		this.queue = this.queue[1:]

		// Enqueue operation
		this.queue = append(this.queue, el)
	}

	// Dequeue operation
	el := this.queue[0]
	this.queue = this.queue[1:]

	return el
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
