package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

const version string = "0.1.0"

var (
	vf = flag.BoolP("version", "v", false, "Print version for rat")
	hf = flag.BoolP("help", "h", false, "Print help")
	lf = flag.StringP("language", "l", "", "Select language to use")
	sf = flag.String("style", "dracula", "Choose chroma style to use")
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

	files := []file{}
	// keep reusing this shit
	var f file
	for _, arg := range args {
		if arg == "-" {
			str, err := readStdin()
			if err != nil {
				log.Error("Unexpected error when reading input from stdin", "err", err)
				return
			}
			f = file{
				content:  strings.Join(str, "\n"),
				filename: "-",
			}
		} else {
			bytes, err := os.ReadFile(arg)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					log.Error("No file exists at path %s", arg)
					return
				}
				log.Error("Error when trying to read file", "err", err)
				return
			}
			f = file{
				content:  string(bytes),
				filename: arg,
			}
		}
		files = append(files, f)
		f.detect()
	}
	fmt.Println(len(files))
}
