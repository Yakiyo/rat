package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	flag "github.com/spf13/pflag"
)

const version string = "0.1.0"

var (
	vf = flag.BoolP("version", "v", false, "Print version for rat")
	hf = flag.BoolP("help", "h", false, "Print help")
)

func init() {
	print := func(a ...any) { fmt.Fprintln(os.Stderr, a...) }
	flag.Usage = func() {
		print("Concatenate FILE(s) to standard output.")
		print("Usage: rat [OPTIONS] [ARGS]")
		print("When no args is provided or `-` is given, stdin is used for input")
		print()
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *vf {
		fmt.Println("Rat version", version)
		return
	}

	if *hf {
		flag.Usage()
		return
	}
	args := flag.Args()
	if len(args) < 1 {
		args = append(args, "-")
	}
	files := []string{}
	for _, arg := range args {
		if arg == "-" {
			str, err := readStdin()
			if err != nil {
				log.Error("Unexpected error when reading input from stdin", "err", err)
				return
			}
			files = append(files, strings.Join(str, "\n"))
			continue
		}
		file, err := os.ReadFile(arg)
		if err != nil {
			log.Error("Error when trying to read file", "err", err)
			return
		}
		files = append(files, string(file))
	}
	fmt.Println(len(files))
}
