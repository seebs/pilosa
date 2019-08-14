package main

import (
	"fmt"

	"github.com/pilosa/pilosa/data"
)

func main() {
	var b = &MyBitmap{}
	var b2 = &MyBitmap{}
	u := data.LookupOpBitmapViewBitmap(b, "Union")
	fmt.Printf("%p\n", u)
	fmt.Printf("%p\n", b.UnionViewBitmap)
	b.UnionViewBitmap(b2)
	u(b2)
	v := data.LookupOpBitmapViewBitmap(b, "Union")
	fmt.Printf("%p\n", v)
}
