# Write preferred mutex

### The problem with a read-preferring mutex
Our initial implementation of the mutex had the potential to cause write-starvation. This can happen when multiple goroutines are basically trading the readers lock among them without giving a chance to a goroutine who wants to acquire the write lock.

### How to implement a write-preferred mutex
To make a mutex write preffered, we just need to make sure that we don't allow any more read locks to be acquired when a write lock is waiting to be acquired. To achieve this, we can track the number of goroutines waiting to acquire a write lock using an `int` property on the mutex.

