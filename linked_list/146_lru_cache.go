package linkedlist

type doubleLinkedNode struct {
	// The reason to store the key is when removing from the map, you need to reference the key
	// You need a doubly linked list because when removing from the middle of the list, it can be an O(1) operation compared to a singly linked list
	key, value int
	prev, next *doubleLinkedNode
}

type LRUCache struct {
	// This maps a key to the address of the double linked list
	nodeMap  map[int]*doubleLinkedNode
	head     *doubleLinkedNode
	tail     *doubleLinkedNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	headNode := &doubleLinkedNode{}
	tailNode := &doubleLinkedNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return LRUCache{
		nodeMap:  make(map[int]*doubleLinkedNode),
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
		lru := this.head.next
		this.removeNode(lru)
		delete(this.nodeMap, lru.key)
	}

	node = &doubleLinkedNode{
		key:   key,
		value: value,
	}
	this.moveToMostRecent(node)
	this.nodeMap[key] = node
}

func (this *LRUCache) removeNode(node *doubleLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToMostRecent(node *doubleLinkedNode) {
	node.next = this.tail
	node.prev = this.tail.prev

	this.tail.prev.next = node
	this.tail.prev = node
}
