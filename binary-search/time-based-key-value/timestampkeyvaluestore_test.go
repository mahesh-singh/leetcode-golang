package timebasedkeyvalue

import "testing"

func TestBSearch(t *testing.T) {
	timestampValues := []timestampValue{
		{timeStamp: 1,
			value: "foo"},
		{timeStamp: 2,
			value: "foo2",
		},
		{timeStamp: 4,
			value: "foo4",
		},
		{timeStamp: 5,
			value: "foo5",
		},
		{timeStamp: 6,
			value: "foo6",
		},
	}

	got := bsearch(timestampValues, -1)
	if got.value != "" {
		t.Errorf("Expected: %v, got: %v", "", got.value)
	}

	got = bsearch(timestampValues, 3)
	if got.value != "foo2" {
		t.Errorf("Expected: %v, got: %v", "foo2", got.value)
	}

	got = bsearch(timestampValues, 5)
	if got.value != "foo5" {
		t.Errorf("Expected: %v, got: %v", "foo5", got.value)
	}

	got = bsearch(timestampValues, 7)
	if got.value != "foo6" {
		t.Errorf("Expected: %v, got: %v", "foo6", got.value)
	}
}

func TestGet(t *testing.T) {
	timeMap := Constructor()
	timeMap.set("foo", "bar", 1)
	got := timeMap.get("foo", 1)
	if got != "bar" {
		t.Fatalf("Expected: %s, got: %s", "bar", got)
	}
	timeMap.get("foo", 3)
	if got != "bar" {
		t.Fatalf("Expected: %s, got: %s", "bar", got)
	}
	timeMap.set("foo", "bar2", 4)

	got = timeMap.get("foo", 4)
	if got != "bar2" {
		t.Fatalf("Expected: %s, got: %s", "bar2", got)
	}
	got = timeMap.get("foo", 5)
	if got != "bar2" {
		t.Fatalf("Expected: %s, got: %s", "bar2", got)
	}
}
