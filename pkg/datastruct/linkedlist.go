package datastruct

//list *listCreate(void)
//void listRelease(list *list)
//list *listAddNodeHead(list *list, void *value)
//list *listAddNodeTail(list *list, void *value)
//list *listInsertNode(list *list, listNode *old_node, void *value, int after)
//void listDelNode(list *list, listNode *node)
//listIter *listGetIterator(list *list, int direction)
//listNode *listNext(listIter *iter)
//void listReleaseIterator(listIter *iter)
//list *listDup(list *orig)
//listNode *listSearchKey(list *list, void *key)
//listNode *listIndex(list *list, long index)
//void listRewind(list *list, listIter *li)
//void listRewindTail(list *list, listIter *li)
//void listRotate(list *list)

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value *interface{}
}

type List struct {
	head  *ListNode
	tail  *ListNode
	len   int
	dup   func()
	free  func()
	match func()
}

func (l *List) ListLength() int {
	return l.len
}

func (l *List) ListFirst() *ListNode {
	return l.head
}
func (l *List) ListLast() *ListNode {
	return l.tail
}

func (l *ListNode) ListPrevNode() *ListNode {
	return l.prev
}

func (l *ListNode) ListNextNode() *ListNode {
	return l.next
}

func (l *ListNode) ListNodeValue() *interface{} {
	return l.value
}

func (l *List) ListSetDupMethod(dupF func()) *List {
	l.dup = dupF
	return l
}

func (l *List) ListSetFreeMethod(freeF func()) *List {
	l.dup = freeF
	return l
}

func (l *List) ListSetMatchMethod(matchF func()) *List {
	l.dup = matchF
	return l
}

func ListCreate() *List {
	list := new(List)
	list.len = 0
	list.head = nil
	list.tail = nil
	return list
}

func (l *List) ListRelease() *List {
	panic("implement me")
	return nil
}

func (l *List) ListAddNodeHead(v *interface{}) *List {
	node := new(ListNode)
	node.value = v
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
	return l
}

func (l *List) ListAddNodeTail(v *interface{}) *List {
	node := new(ListNode)
	node.value = v
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
	return l
}

func (l *List) ListInsertNode(oldNode *ListNode, v *interface{}, after bool) *List {
	node := new(ListNode)
	node.value = v

	if after {
		if oldNode == l.tail { // oldNode.next != nil
			l.tail = node
		}
		node.next = oldNode.next
		node.prev = oldNode
	} else {
		if oldNode == l.head { //oldNode.prev != nil
			l.head = node
		}
		node.prev = oldNode.prev
		node.next = oldNode
	}

	if node.prev != nil {
		node.prev.next = node
	}
	if node.next != nil {
		node.next.prev = node
	}

	l.len++
	return nil
}

func (l *List) ListDelNode(oldNode *ListNode) *List {
	if l.head == oldNode {
		l.head = oldNode.next
	} else {
		oldNode.prev.next = oldNode.next
	}

	if l.tail == oldNode {
		l.tail = oldNode.prev
	} else {
		oldNode.next.prev = oldNode.prev
	}
	oldNode.prev = nil
	oldNode.next = nil
	l.len--
	return l
}

func (l *List) ListGetIterator() *List {
	return nil
}

func (l *List) ListNext() *List {
	return nil
}

func (l *List) ListReleaseIterator() *List {
	return nil
}

func (l *List) ListDup() *List {
	return nil
}

func (l *List) ListSearchKey() *List {
	return nil
}

func (l *List) ListIndex() *List {
	return nil
}

func (l *List) ListRewind() *List {
	return nil
}

func (l *List) ListRewindTail() *List {
	return nil
}

func (l *List) ListRotate() *List {
	return nil
}
