# goint

Generic pOINTy

Inspired by [openlyinc/pointy](https://github.com/openlyinc/pointy).
Provides a one generic function for getting pointers to variables
and another one for dereferencing pointers.

Two Functions to rule them all

> NOTE: Requires go 1.18+

# Installation
```
go get github.com/majakubiec/goint
```

# Example
```golang
package main

import (
	"fmt"

	"github.com/majakubiec/goint"
)

func main() {
	foo := goint.To(2022)
	fmt.Println("foo is a pointer to:", *foo)

	bar := goint.To("point to me")
	fmt.Println("bar is a pointer to:", *bar)

	// get the value back out
	barVal := goint.From(bar, "empty!")
	fmt.Println("bar's value is:", barVal)
}
```