package main

type DLLNode struct {
	key, value int
	prev, next *DLLNode
}

type LRUCache struct {
	capacity int
	cache map[int]*DLLNode
	head, tail *DLLNode
}

func (this *LRUCache) Constructor(capacity int) *LRUCache {
	l := &LRUCache{
		capacity: capacity,
		cache: make(map[int]*DLLNode),
		head: &DLLNode{},
		tail: &DLLNode{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return  node.value
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &DLLNode{key: key, value: value}
		this.cache[key] = newNode
		// if we've reached capacity, remove lru from tail
		if len(this.cache) > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)			
		}
	}
}

func (this *LRUCache) removeNode(node *DLLNode){
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLLNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node 
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLLNode){
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLLNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}