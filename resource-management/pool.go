package main

import (
	"fmt"
	"sync"
)

func main() {
	// this creates and allocate a new set of objects of 1kb
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("allocate new objects")
			return make([]byte, 1024)
		},
	}

	// We use the get method for getting an obj from the pool
	obj := pool.Get().([]byte)
	fmt.Printf("The lenght of the obj is: %d\n", len(obj))

	// The obj is returned to the pool
	pool.Put(obj)

	// instead of allocating a new obj, we get the already created
	sameobj := pool.Get().([]byte)
	fmt.Printf("The same obj returned: %d\n", len(sameobj))

	pool.Put(sameobj)
}
