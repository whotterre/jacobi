package main
// Implementation of a hash table with seperate chaining for collision resolution.
import (
	"hash/fnv"
)

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type HashTable struct {
	size  int
	table []*Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*Node, size),
	}
}
// Generates a hash using FNV-1a algorithm
func (ht *HashTable) hash(key string) int {
	h := fnv.New32a()
	// Convert the key to bytes and write it to the hash
	h.Write([]byte(key))
	return int(h.Sum32()) % ht.size
}

// Returns the size of the hash table
func (ht *HashTable) Size() int {
	return ht.size
}

// Inserts a key-value pair into the hash table.
func (ht *HashTable) Insert(key string, value interface{}) {
	index := ht.hash(key)
	newNode := &Node{key: key, value: value}

	if ht.table[index] == nil {
		ht.table[index] = newNode
	} else {
		current := ht.table[index]
		for current.next != nil {
			if current.key == key { 
				current.value = value
				return
			}
			current = current.next
		}
		current.next = newNode
	}
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.hash(key)
	current := ht.table[index]
    // Traverse the linked list at the index
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return nil, false
}

