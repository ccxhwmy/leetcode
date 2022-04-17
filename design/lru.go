package design

//https://leetcode.com/problems/lru-cache

type Node struct {
	key        int
	val        int
	prev, next *Node
}

type LRUCache struct {
	limit int
	cache map[int]*Node
	dummyHead, dummyTail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.next = head
	return LRUCache{
		limit:     capacity,
		cache:     make(map[int]*Node, 0),
		dummyHead: head,
		dummyTail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	if this.limit <= len(this.cache) {
		delKey := this.removeLast()
		delete(this.cache, delKey)
	}
	addNode := &Node{
		key:  key,
		val:  value,
	}
	this.cache[key] = addNode
	this.addFirst(addNode)
}

func (this *LRUCache) moveToHead(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	this.addFirst(node)
}

func (this *LRUCache) addFirst(node *Node) {
	node.next = this.dummyHead.next
	node.prev = this.dummyHead
	node.prev.next = node
	node.next.prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) removeLast() int {
	lastNode := this.dummyTail.prev
	this.remove(lastNode)
	return lastNode.key
}
