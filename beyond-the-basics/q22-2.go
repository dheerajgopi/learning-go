package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader, lineNumFlag bool) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if lineNumFlag {
			fmt.Fprintf(os.Stdout, "%d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	lineNumFlag := flag.Bool("n", false, "show line number")
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin), *lineNumFlag)
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(file), *lineNumFlag)
	}
}
