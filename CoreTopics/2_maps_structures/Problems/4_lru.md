LRU Cache:

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
Example: 
//create cache of size 2
cache := Constructor(2)
cache.Put(1, v1)
cache.Put(2, v2)
cache.Get(1)    // returns v1
cache.Put(3, v3) // removes key 2
cache.Get(2)    // returns -1 (not found)