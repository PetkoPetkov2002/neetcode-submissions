type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	keyMap   map[int]*Node
	headNode *Node
	tailNode *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyMap:   make(map[int]*Node),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.keyMap[key]
	if !ok {
		return -1
	}

	this.moveToTail(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	if node, ok := this.keyMap[key]; ok {
		node.value = value
		this.moveToTail(node)
		return
	}

	if len(this.keyMap) == this.capacity {
		oldest := this.headNode
		this.removeNode(oldest)
		delete(this.keyMap, oldest.key)
	}

	node := &Node{
		key:   key,
		value: value,
	}

	this.keyMap[key] = node
	this.appendToTail(node)
}

// removeNode assumes node belongs to this list.
func (this *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.headNode = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tailNode = node.prev
	}

	node.prev = nil
	node.next = nil
}

// appendToTail assumes node is detached from any list.
func (this *LRUCache) appendToTail(node *Node) {
	node.prev = this.tailNode
	node.next = nil

	if this.tailNode != nil {
		this.tailNode.next = node
	} else {
		this.headNode = node
	}

	this.tailNode = node
}

func (this *LRUCache) moveToTail(node *Node) {
	if node == this.tailNode {
		return
	}

	this.removeNode(node)
	this.appendToTail(node)
}