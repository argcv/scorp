package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/argcv/scorp/pkg/cli"
)

func main() {
	cli.ReadLines(os.Stdin, func(buf string) bool {
		if i, e := strconv.ParseInt(buf, 2, strconv.IntSize); e != nil {
			fmt.Println("-")
		} else {
			fmt.Printf("0x%x\n", i)
		}
		return true
	})
}
