package lru

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.

type Cache struct {
	ll       *list.List		//在这里我们直接使用 Go 语言标准库实现的双向链表list.List。
	cache    map[string]*list.Element	//键是字符串，值是双向链表中对应节点的指针。

	maxBytes int64	//最大内存
	nBytes   int64	//当前已使用的内存

	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)	//某条记录被移除时的回调函数，可以为 nil
}
//键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}
