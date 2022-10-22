package lrucache

import "testing"

/*
["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],		[1,1],[2,2],[1],   [3,3],[2],  [4,4], [1], [3],  [4]]

* Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
*/

func TestLRUCache(t *testing.T) {
	lruC := Constructor(2)
	lruC.Put(1, 1)
	lruC.Put(2, 2)
	out := lruC.Get(1)
	if out != 1 {
		t.Errorf("Expected: %v, out %v", 1, out)
	}
	lruC.Put(3, 3)
	out = lruC.Get(2)
	if out != -1 {
		t.Errorf("Expected: %v, out %v", -1, out)
	}
	lruC.Put(4, 4)
	out = lruC.Get(1)
	if out != -1 {
		t.Errorf("Expected: %v, out %v", -1, out)
	}
	out = lruC.Get(3)
	if out != 3 {
		t.Errorf("Expected: %v, out %v", 3, out)
	}

	out = lruC.Get(4)
	if out != 4 {
		t.Errorf("Expected: %v, out %v", 4, out)
	}

}
