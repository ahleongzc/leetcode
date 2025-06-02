package stack

type MinStack struct {
	normalStack []int
	minStack    []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		normalStack: make([]int, 0),
		minStack:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.normalStack = append(this.normalStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
		return
	}

	currentMin := this.minStack[len(this.minStack)-1]
	currentMin = min(currentMin, val)
	this.minStack = append(this.minStack, currentMin)
}

func (this *MinStack) Pop() {
	this.normalStack = this.normalStack[:len(this.normalStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.normalStack[len(this.minStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
