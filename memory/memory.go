package memory

import "sync"

type Memory struct {
	mutex sync.Mutex
	Mem   map[string]string
	Limit int
	list  LinkedList
}

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewMemory(limit int) Memory {
	return Memory{Mem: make(map[string]string), list: NewLinkedList(""), Limit: limit}
}

func NewLinkedList(x string) LinkedList {
	node := &Node{x, nil}

	return LinkedList{node, 1}
}

// This function belongs to the Linked List object that contains a linked list and performs the operation of adding
// a new element to the beginning of the list. First, the current pointer is set to head, then a temporary
// var tmp is set which will contain a new node with the received value and next set to nil. Then, tmp is assigned
// as the next from the last item in the list. Finally, the size of the list is increased by one.

func (self *LinkedList) push(value string) {

	current := self.Head
	var tmp *Node

	for current.Next != nil {
		current = current.Next
	}

	tmp = &Node{Value: value, Next: nil}

	current.Next = tmp

	self.Size += 1
}

// This function belongs to the linked list object and is responsible for removing the first item from the list.

func (self *LinkedList) Delete() {

	self.Head = self.Head.Next

	self.Size -= 1
}

// The Get function is part of the Memory object and is responsible for retrieving the value corresponding to a
// certain key. Before accessing the contents of shared memory, the functionality waits for other users who may be
// modifying the data structure. Then, it checks if the indicated key exists in the memory table, and
// if so, returns its value. Otherwise, it returns a string indicating that the value was not found.

func (self Memory) Get(key string) string {

	self.mutex.Lock()
	defer self.mutex.Unlock()

	if value, ok := self.Mem[key]; ok {
		return value
	} else {
		return "not found"
	}
}

// The Load function belongs to the Memory object and is used to load information into shared memory.
// It is passed two arguments, key and info. Previously, it checks if the number of items in the
// list exceeds the set limit. If this is true, it removes the last item from the list. Then, it adds
// the element provided in the key argument to the list and stores the info parameter next to it.

func (self Memory) Load(key string, info string) {

	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.list.Size > self.Limit {
		self.list.Delete()
	}

	self.list.push(key)
	self.Mem[key] = info
}
