// 输出参数的索引和值，每行一个
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}
}
