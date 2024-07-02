Geecache is a distributed cache system.

## What's the problems of distributed cache system?
1. What if RAM is not enough for data?
We must delete some data to spare up space for new data? But what data shall we delete? We could delete some data by random or in time order(FIFO) or based on the frequence the data used(LRU)
2. We must consider the behaviour in concurrency environment. Access to the cache is generally impossible to be serial. Map does not have concurrency protection. To cope with concurrency scenarios, modification operations (including addition, update and deletion) need to be locked.
3. What if the performance of a single server is not enough
The resources of a single server are limited. With the increase in business volume and traffic, a single server is not able to cope with traffic. So maybe distributed system is a good idea to expand the performance of the system, which is so called Horizontal expansion

## lru.go
LRU mainly implement the based data structure
  +----+    +----+    +----+    +----+    +----+
  | k1 |    | k2 |    | k3 |    | k4 |    | k5 |
  +----+    +----+    +----+    +----+    +----+
    |         |         |         |         |
    v         v         v         v         v
+----+ <-> +----+ <-> +----+ <-> +----+ <-> +----+
| n1 | <-> | n2 | <-> | n3 | <-> | n4 | <-> | n5 |
+----+     +----+     +----+     +----+     +----+
Head                                         End

- k.. are map, storing key and value. 
- n.. are double linked listed. And the element in double linked listed is as following. The advantage of storing key in linked list is that when delete the element in linked list, we can use the key stored in linked list to delte the kv in map
```
type entry struct {
	key   string
	value Value
}
```
## cache.go
cache is a wrapper of the lru, with added locks to achieve concurrency protection
