package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func writeToBuffer(data string) {
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)

	buffer.Reset()

	buffer.WriteString(data)
	fmt.Println(buffer.String())
	fmt.Printf("Buffers capacity: %d\n", buffer.Cap())
	fmt.Printf("Buffers length: %d\n", buffer.Len())
}

func main() {
	writeToBuffer("This is buffers in a pool1")
	writeToBuffer("This is another buffer in a pool2")
	writeToBuffer("This is one more buffer in a pool")
}
