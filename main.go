package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

const version string = "0.1.0"

var (
	vf = pflag.BoolP("version", "v", false, "Print version for rat")
	hf = pflag.BoolP("help", "h", false, "Print help")
	lf = pflag.StringP("language", "l", "", "Select language to use")
	sf = pflag.String("style", "dracula", "Choose chroma style to use")
)

func run() error {
	if *vf {
		fmt.Println("Rat version", version)
		return nil
	}

	if *hf {
		pflag.Usage()
		return nil
	}
	args := pflag.Args()
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
				// log.Error("Unexpected error when reading input from stdin", "err", err)
				err = errors.Join(errors.New("unexpected error when reading input from stdin"), err)
				return err
			}
			f = file{
				content:  strings.Join(str, "\n"),
				filename: "-",
			}
		} else {
			bytes, err := os.ReadFile(arg)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Errorf("no file exists at path %s", arg)

				}
				return errors.Join(errors.New("error when trying to read file"), err)
			}
			f = file{
				content:  string(bytes),
				filename: arg,
			}
		}
		files = append(files, f)
	}
	_, noc := os.LookupEnv("NO_COLOR")
	for _, file := range files {
		log.Info("printing file", "filename", file.filename)
		if *sf == "none" || noc {
			fmt.Println(file.content)
			continue
		}
		err := file.format()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	initLogger()

	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Concatenate FILE(s) to standard output.")
		fmt.Fprintln(os.Stderr, "Usage: rat [OPTIONS] [ARGS]")
		fmt.Fprintln(os.Stderr, "When no args is provided or `-` is given, stdin is used for input")
		fmt.Fprintln(os.Stderr)
		pflag.PrintDefaults()
	}

	pflag.Parse()

	err := run()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
