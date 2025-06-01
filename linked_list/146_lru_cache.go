package linkedlist

type Node struct {
	// The reason to store the key is when removing from the map, you need to reference the key
	key, value int
	prev, next *Node
}

type LRUCache struct {
	// This maps a key to the address of the double linked list
	nodeMap  map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	headNode := &Node{}
	tailNode := &Node{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return LRUCache{
		nodeMap:  make(map[int]*Node),
		head:     headNode,
		tail:     tailNode,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.nodeMap[key]
	if !exists {
		return -1
	}
	this.removeNode(node)
	this.moveToMostRecent(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.nodeMap[key]
	if exists {
		node.value = value
		this.removeNode(node)
		this.moveToMostRecent(node)
		return
	}

	// Evict the least recently used
	if len(this.nodeMap) == this.capacity {
		LRU := this.head.next
		this.removeNode(LRU)
		delete(this.nodeMap, LRU.key)
	}

	node = &Node{
		key:   key,
		value: value,
	}
	this.moveToMostRecent(node)
	this.nodeMap[key] = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToMostRecent(node *Node) {
	node.next = this.tail
	node.prev = this.tail.prev

	this.tail.prev.next = node
	this.tail.prev = node
}
