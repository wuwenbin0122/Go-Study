// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var seq = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *seq))
	if !*n {
		fmt.Println()
	}
}
