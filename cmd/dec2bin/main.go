package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/argcv/scorp/pkg/cli"
)

func main() {
	cli.ReadLines(os.Stdin, func(buf string) bool {
		if i, e := strconv.ParseInt(buf, 10, strconv.IntSize); e != nil {
			fmt.Println("-")
		} else {
			fmt.Printf("%0b\n", i)
		}
		return true
	})
}
