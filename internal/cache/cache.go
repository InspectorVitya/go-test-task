package cache

import (
	"errors"
	"fmt"
	"sync"
)

type Key string

var ErrCache = errors.New("cannot execute NewCache: capacity must be greater than 0")

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
	mx       *sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mx.Lock()
	defer l.mx.Unlock()

	item, ok := l.items[key]
	if !ok {
		newItem := cacheItem{
			key:   key,
			value: value,
		}
		itemList := l.queue.PushFront(newItem)
		l.items[key] = itemList
		if l.queue.Len() > l.capacity {
			lastItem := l.queue.Back()
			l.queue.Remove(lastItem)
			delete(l.items, lastItem.Value.(cacheItem).key)
		}
		return false
	}
	item.Value = cacheItem{
		key:   key,
		value: value,
	}
	l.queue.PushFront(item)
	return true
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mx.Lock()
	defer l.mx.Unlock()

	item, ok := l.items[key]
	if !ok {
		return nil, false
	}
	l.queue.MoveToFront(item)
	return item.Value.(cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.capacity = 0
	l.queue = &list{}
	l.items = make(map[Key]*listItem)
}

func (l *lruCache) Test(key Key) {
	for val, key := range l.items {
		fmt.Println(val, key)
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) (Cache, error) {
	if capacity > 0 {
		return &lruCache{
			capacity: capacity,
			queue:    &list{},
			items:    make(map[Key]*listItem),
			mx:       &sync.Mutex{},
		}, nil
	}
	return nil, ErrCache
}
