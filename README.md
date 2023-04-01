# Ring

![test](https://github.com/WinPooh32/ring/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/ring.svg)](https://pkg.go.dev/github.com/WinPooh32/ring)

Generic circular buffer

## Example

```Go
package main

import (
	"fmt"

	"github.com/WinPooh32/ring"
)

func main() {
	const capacity = 9

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	ringBuf := ring.Make[int](capacity)

	for i, v := range data {
		pop, elem := ringBuf.Push(v)
		if pop {
			fmt.Println(elem)
		}
	}

	l, r = ringBuf.TwoParts()

	fmt.Println(l, r)
}
```
Output:
```
0
1
2
3
4
5
6
7
[8] [9 10 11 12 13 14 15 16]
```