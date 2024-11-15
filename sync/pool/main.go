package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("initalising buffer pool")
		return new(bytes.Buffer)
	},
}

func logger(w io.Writer, debug string) {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString("testing")
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufferPool.Put(b)

}

func main() {

	logger(os.Stdout, "debug-1")
	logger(os.Stdout, "debug-2")
	logger(os.Stdout, "debug-3")

}
