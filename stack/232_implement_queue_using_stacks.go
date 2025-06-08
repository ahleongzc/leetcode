package stack

type MyQueue struct {
	enqueueStack []int
	dequeueStack []int
}

func ConstructorQueueUsingStack() MyQueue {
	return MyQueue{
		enqueueStack: make([]int, 0),
		dequeueStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.enqueueStack = append(this.enqueueStack, x)
}

func (this *MyQueue) Pop() int {
	element := this.Peek()
	this.dequeueStack = this.dequeueStack[:len(this.dequeueStack)-1]

	return element
}

func (this *MyQueue) Peek() int {
	if len(this.dequeueStack) == 0 {
		for range len(this.enqueueStack) {
			this.dequeueStack = append(this.dequeueStack, this.enqueueStack[len(this.enqueueStack)-1])
			this.enqueueStack = this.enqueueStack[:len(this.enqueueStack)-1]
		}
	}
	element := this.dequeueStack[len(this.dequeueStack)-1]

	return element
}

func (this *MyQueue) Empty() bool {
	return len(this.dequeueStack) == 0 && len(this.enqueueStack) == 0
}
