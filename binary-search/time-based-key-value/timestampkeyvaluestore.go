package timebasedkeyvalue

/*
981. Time Based Key-Value Store
https://leetcode.com/problems/time-based-key-value-store/
Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".


Example 1:

Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"
*/

/*
Possible solution
1. We can have a map with key as a string and value as an array of timestamp and value pair
2. key: foo, value: [[1, bar], [4, bar2]]
3. Set will put the pair at the end of the array, with the assumption that every new timestamp will have greater than previous one
4. Get request will first filter via key. To search value via timestamp, it should use binary search as they are sorted by timestamp
*/

type timestampValue struct {
	timeStamp int
	value     string
}

type timeMap struct {
	hashMap map[string][]timestampValue
}

func Constructor() timeMap {
	return timeMap{hashMap: make(map[string][]timestampValue)}
}

func (tm *timeMap) set(key string, value string, timeStamp int) {
	if _, ok := tm.hashMap[key]; ok {
		tm.hashMap[key] = append(tm.hashMap[key], timestampValue{timeStamp, value})
	} else {
		tm.hashMap[key] = []timestampValue{{timeStamp, value}}
	}
}

func (tm *timeMap) get(key string, timeStamp int) string {
	if _, ok := tm.hashMap[key]; ok {
		value := bsearch(tm.hashMap[key], timeStamp)
		return value.value
	} else {
		return ""
	}

}

func bsearch(timestampValues []timestampValue, timeStamp int) timestampValue {
	l, r := 0, len(timestampValues)

	for l < r {
		//find mid
		m := (l + r) / 2

		if timestampValues[m].timeStamp <= timeStamp {
			//go right
			l = m + 1
		} else {
			//go left
			r = m
		}
	}

	if r == 0 {
		return timestampValue{-1, ""}
	}
	return timestampValues[r-1]
}
