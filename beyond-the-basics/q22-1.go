package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	if len(os.Args[1:]) == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 1; i < len(os.Args[:]); i++ {
		file, err := os.Open(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], os.Args[i], err.Error())
			continue
		}
		cat(bufio.NewReader(file))
	}
}
