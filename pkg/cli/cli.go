package cli

import (
	"bufio"
	"io"
	"strings"

	"github.com/argcv/stork/log"
)

func ReadLines(rd io.Reader, f func(line string) bool) {
	reader := bufio.NewReader(rd)
	for ; ; {
		buf, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Infof("unexpected error :%v", err)
			}
			break
		}
		buf = strings.TrimSpace(buf)
		if !f(buf) {
			break
		}
	}
}
