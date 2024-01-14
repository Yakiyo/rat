package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

const version string = "0.1.0"

var (
	vf = pflag.BoolP("version", "v", false, "Print version for rat")
	hf = pflag.BoolP("help", "h", false, "Print help")
	lf = pflag.StringP("language", "l", "", "Select language to use")
	sf = pflag.String("style", "dracula", "Choose chroma style to use")
)

func init() {
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Concatenate FILE(s) to standard output.")
		fmt.Fprintln(os.Stderr, "Usage: rat [OPTIONS] [ARGS]")
		fmt.Fprintln(os.Stderr, "When no args is provided or `-` is given, stdin is used for input")
		fmt.Fprintln(os.Stderr)
		pflag.PrintDefaults()
	}
}

func main() {
	pflag.Parse()

	if *vf {
		fmt.Println("Rat version", version)
		return
	}

	if *hf {
		pflag.Usage()
		return
	}
	args := pflag.Args()
	if len(args) < 1 {
		args = append(args, "-")
	}
	if err := setDefaults(); err != nil {
		log.Error(err)
		return
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
	}
	log.Info("finished reading files", "len", len(files))
	for _, file := range files {
		log.Info(file.filename)
		// we prolly don't want pretty printing
		if style == nil {
			fmt.Println(file.content)
			continue
		}
		err := f.detect()
		if err != nil {
			log.Error(err)
			return
		}
		err = f.format()
		if err != nil {
			log.Error(err)
			return
		}
	}
}
