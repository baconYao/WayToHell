package main

func main() {
	lru := &lru{}
	cache := initCache(lru)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("d", "4")

	lfu := &lfu{}
	cache.setEvictionAlgo(lfu)
	cache.add("f", "5")
}
