package hash

const ArraySize = 7

//Hash Table structure
type HashTable struct {
	array [ArraySize]*bucket
}

//Bucket structure
type bucket struct {
	head *bucketNode
}

//Bucket node structure -> linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert -> Add a key to the hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

//insert -> add a key to the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		return
	}
}

//search -> find the key in the bucket and return true or false
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete -> delete a key in the bucket
func (b *bucket) delete(k string) {
	if b.head == nil {
		return
	}
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for currentNode := b.head.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.key == k {
			previousNode.next = currentNode.next
			return
		}
		previousNode = currentNode
	}
}

//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//Init -> Create a bucket in each slot
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}
