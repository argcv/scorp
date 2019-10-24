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
		if len(buf) > 2 && strings.HasPrefix(buf, "0x") {
			buf = buf[2:]
		}
		if i, e := strconv.ParseInt(buf, 16, strconv.IntSize); e != nil {
			fmt.Println("-")
		} else {
			fmt.Println(i)
		}
		return true
	})
}
