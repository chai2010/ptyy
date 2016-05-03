# 莆田系医院查询

[![GoDoc](https://godoc.org/github.com/chai2010/ptyy?status.svg)](https://godoc.org/github.com/chai2010/ptyy)

Install
=======

1. `go get github.com/chai2010/ptyy`
2. `go run hello.go`

Example
=======

```Go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chai2010/ptyy"
)

var (
	flagLimits = flag.Int("limits", 0, "设置最大查询数目")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [query...]\n", filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			for _, v := range ptyy.Search(flag.Arg(i), *flagLimits) {
				fmt.Printf("%s (%s)\n", v.Name, v.City)
			}
		}
	}
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
