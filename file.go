package main

import (
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

type file struct {
	content  string
	filename string

	lexer chroma.Lexer
	style chroma.Lexer
}

func (f *file) detect() error {
	// if lang != "" {
	// 	f.lexer = lexers.Get(lang)
	// 	if f.lexer == nil {
	// 		return errors.New("Unknown language " + lang)
	// 	}
	// 	return nil
	// }
	if lexer != nil {
		return nil
	}
	f.lexer = lexers.Match(f.filename)
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

func (f *file) setStyle(s string) error {
	return nil
}
