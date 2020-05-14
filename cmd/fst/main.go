package main

import (
	"bytes"
	"fmt"

	"github.com/bradleyjkemp/memviz"

	"github.com/eliias/algorithms/pkg/fst"
)

func main() {
	// ordered set, value is slice index
	dict := []string{
		"monday",
		"saturday",
		"thursday",
		"tuesday",
	}

	ds := fst.NewFST()
	for value, word := range dict {
		ds.Add(word, value)
	}
	ds.GC()

	buf := &bytes.Buffer{}
	memviz.Map(buf, ds)
	fmt.Println(buf.String())
}
