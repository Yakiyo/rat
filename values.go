package main

import (
	"errors"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/mattn/go-isatty"
)

var (
	lexer     chroma.Lexer
	style     *chroma.Style
	formatter = func() chroma.Formatter { return formatters.Get("terminal16m") }
)

func setDefaults() error {
	lang := *lf
	s := *sf
	if lang != "" {
		lexer = lexers.Get(lang)
		if lexer == nil {
			return errors.New("Unknown language " + lang)
		}
		return nil
	}
	if s == "none" || !isAtty() {
		style = nil
		return nil
	}
	style = styles.Get(s)
	if style == nil {
		return errors.New("Unknown style " + s)
	}
	return nil
}

func isAtty() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}
