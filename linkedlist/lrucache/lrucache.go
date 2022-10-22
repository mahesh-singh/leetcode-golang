/*
146. LRU Cache
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.



Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
*/

package lrucache

type LRUCache struct {
	Capacity int
	Cache    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Capacity: capacity, Cache: make(map[int]*node)}
}

func (lruCache *LRUCache) Get(key int) int {
	if currNode, ok := lruCache.Cache[key]; ok {
		lruCache.remove(currNode)
		lruCache.add(currNode)
		return currNode.val
	}
	return -1
}

func (lruCache *LRUCache) Put(key int, value int) {
	if currNode, ok := lruCache.Cache[key]; ok {
		currNode.val = value
		lruCache.Cache[key] = currNode
	} else {
		n := &node{key: key, val: value}
		lruCache.add(n)
		lruCache.Cache[key] = n
	}

	if len(lruCache.Cache) > lruCache.Capacity {
		delete(lruCache.Cache, lruCache.tail.key)
		lruCache.remove(lruCache.tail)
	}
}
func (lruCache *LRUCache) add(n *node) {
	n.prev = nil
	n.next = lruCache.head
	if lruCache.head != nil {
		lruCache.head.prev = n
	}
	lruCache.head = n
	if lruCache.tail == nil {
		lruCache.tail = lruCache.head
	}
}

func (lruCache *LRUCache) remove(n *node) {
	//if node is tail
	if lruCache.tail == n {
		lruCache.tail = n.prev
		lruCache.tail.next = nil
	} else if lruCache.head == n {
		//if node is head
		lruCache.head = n.next
		lruCache.head.prev = nil
	} else {
		n.prev.next = n.next.prev
		n.next.prev = n.prev.next
	}

}
