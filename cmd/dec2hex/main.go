package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/argcv/scorp/pkg/cli"
)

func main() {
	cli.ReadLines(os.Stdin, func(buf string) bool {
		buf = strings.ToLower(buf)
		if i, e := strconv.ParseInt(buf, 10, strconv.IntSize); e != nil {
			fmt.Println("-")
		} else {
			fmt.Printf("0x%x\n", i)
		}
		return true
	})
}
