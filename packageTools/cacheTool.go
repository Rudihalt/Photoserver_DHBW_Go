package packageTools

// Qnode ...
type Qnode struct {
	key int
	value string
	prev, next *Qnode
}

func addQNode(key int, value string) *Qnode {
	return &Qnode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

// Queue ...
type Queue struct {
	front *Qnode
	rear  *Qnode
}

func (q *Queue) isEmpty() bool {
	return q.rear == nil
}

func (q *Queue) addFrontPage(key int, value string) *Qnode {
	page := addQNode(key, value)
	if q.front == nil && q.rear == nil {
		q.front, q.rear = page, page
	} else {
		page.next = q.front
		q.front.prev = page
		q.front = page
	}
	return page
}

func (q *Queue) moveToFront(page *Qnode) {
	if page == q.front {
		return
	} else if page == q.rear {
		q.rear = q.rear.prev
		q.rear.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	page.next = q.front
	q.front.prev = page
	q.front = page
}

func (q *Queue) removeRear() {
	if q.isEmpty() {
		return
	} else if q.front == q.rear {
		q.front, q.rear = nil, nil
	} else {
		q.rear = q.rear.prev
		q.rear.next = nil
	}
}

func (q *Queue) getRear() *Qnode {
	return q.rear
}

// LRUCache ...
type LRUCache struct {
	capacity, size int
	pageList       Queue
	pageMap        map[int]*Qnode
}

func (lru *LRUCache) InitLru(capacity int) {
	lru.capacity = capacity
	lru.pageMap = make(map[int]*Qnode)
}

func (lru *LRUCache) Get(key int) string {
	if _, found := lru.pageMap[key]; !found {
		return ""
	}
	val := lru.pageMap[key].value
	lru.pageList.moveToFront(lru.pageMap[key])
	return val
}

func (lru *LRUCache) Put(key int, value string) {
	if _, found := lru.pageMap[key]; found {
		lru.pageMap[key].value = value
		lru.pageList.moveToFront(lru.pageMap[key])
		return
	}

	if lru.size == lru.capacity {
		key := lru.pageList.getRear().key
		lru.pageList.removeRear()
		lru.size--
		delete(lru.pageMap, key)
	}
	page := lru.pageList.addFrontPage(key, value)
	lru.size++
	lru.pageMap[key] = page
}

