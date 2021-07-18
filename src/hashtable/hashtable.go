package hashtable

import (
	"sync"
)

// A hash table associates a set of keys with a set of values.
// Each key/value pair is an entry in the table.
// Given a key, you can look up its corresponding value.
// You can add new key/value pairs and remove entries by key.
// If you add a new value for an existing key, it replaces the previous entry.
// The average time complexity is O(1).
// The average space complexity is O(n).
//
// Each key/value pair is called an entry,
// we will the entries in an array.
// Each position in the array is called a bucket.
//
// Techniques used to build a hash table:
//
// Separate chaning
// With this technique, each bucket contains a linked list of entries,
// if a collision happens we just append the entry to the linked list.
// It's easy to implement but it has a lot of overhead from pointers
// and it is not cache friendly.
//
// Open addressing
// With this technique, all entries live direcly in the bucket array,
// with one entry per bucket. If two entries collide in the same
// bucket, we find a different empty bucket to use instead.
// It's more complex but cache friendly.
type entry struct {
	key   string
	value interface{}
}

var tombstone = &entry{
	key:   "tombstone",
	value: nil,
}

func (e *entry) IsEmpty() bool {
	return e == nil || e == tombstone
}

type hashtable struct {
	mutex    sync.RWMutex
	size     int
	capacity int
	buckets  []*entry
}

type T = hashtable

// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash
func hash(s string) int {
	var out int = 2166136261

	for i := 0; i < len(s); i++ {
		out ^= int(s[i])
		out *= 16777619
	}

	return out
}

// Table capacity is exceeded if buckets are at least 75% full.
func (table *hashtable) willExceedCapacity() bool {
	const hashtableMaxLoad = 0.75

	return table.size+1 > int(float64(table.capacity)*hashtableMaxLoad)
}

func (table *hashtable) doubleCapacity() {
	table.mutex.Lock()
	defer table.mutex.Unlock()

	buckets := table.buckets

	table.size = 0
	table.capacity = 2*table.capacity + 1
	table.buckets = make([]*entry, table.capacity)

	for _, bucket := range buckets {
		if bucket.IsEmpty() {
			continue
		}

		table.insert(bucket.key, bucket.value)
	}
}

func (table *hashtable) IsEmpty() bool {
	return table.Size() == 0
}

func (table *hashtable) find(key string) *entry {
	if table.IsEmpty() {
		return nil
	}

	index := hash(key) % table.capacity

	visitedBuckets := 0

	for {
		if visitedBuckets == table.Size() {
			return nil
		}

		entry := table.buckets[index]
		if entry == nil {
			return nil
		}

		if entry.key == key {
			return entry
		}

		visitedBuckets++
		index = (index + 1) % table.capacity
	}
}

func (table *hashtable) Has(key string) bool {
	table.mutex.RLock()
	defer table.mutex.RUnlock()

	return table.find(key) != nil
}

func (table *hashtable) Get(key string) interface{} {
	table.mutex.RLock()
	defer table.mutex.RUnlock()

	entry := table.find(key)
	if entry == nil {
		return nil
	}

	return entry.value
}

func (table *hashtable) Remove(key string) {
	if table.IsEmpty() {
		return
	}

	table.mutex.Lock()
	defer table.mutex.Unlock()

	index := hash(key) % table.capacity

	visitedBuckets := 0

	for {
		if visitedBuckets == table.Size() {
			return
		}

		entry := table.buckets[index]
		if entry == nil {
			return
		}

		if entry.key == key {
			table.buckets[index] = tombstone
			table.size--
			return
		}

		visitedBuckets++
		index = (index + 1) % table.capacity
	}
}

func (table *hashtable) insert(key string, value interface{}) {
	index := hash(key) % table.capacity

	for {
		if table.buckets[index] == nil {
			table.buckets[index] = &entry{
				key,
				value,
			}

			table.size++
			return
		} else if table.buckets[index].key == key {
			table.buckets[index].value = value
			return
		}

		index = (index + 1) % table.capacity
	}
}

func (table *hashtable) Set(key string, value interface{}) {
	if table.willExceedCapacity() {
		table.doubleCapacity()
	}

	table.mutex.Lock()
	defer table.mutex.Unlock()

	table.insert(key, value)
}

func (table *hashtable) Size() int {
	return table.size
}

func New() T {
	return T{
		mutex:    sync.RWMutex{},
		size:     0,
		capacity: 0,
		buckets:  nil,
	}
}
