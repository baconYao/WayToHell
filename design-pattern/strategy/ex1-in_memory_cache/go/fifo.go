package main

import "fmt"

type fifo struct {
}

func (f *fifo) evict(c *cache) {
	fmt.Println("Evicting by FIFO algorithm")
}
