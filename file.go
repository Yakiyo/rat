package main

import (
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

type file struct {
	content  string
	filename string

	lexer chroma.Lexer
}

func (f *file) detect() error {
	if lexer != nil {
		return nil
	}
	if f.filename != "-" {
		f.lexer = lexers.Match(f.filename)
	}
	// if we can't match by filename, try analyzing content
	if f.lexer == nil {
		f.lexer = lexers.Analyse(f.content)
	}
	// even if that doesn't work, switch to fallback
	if f.lexer == nil {
		f.lexer = lexers.Fallback
	}
	return nil
}

func (f *file) format() error {
	iter, err := f.lexer.Tokenise(nil, f.content)
	if err != nil {
		return err
	}
	err = formatter().Format(os.Stdout, style, iter)
	return err
}
