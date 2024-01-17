package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/log"
	"github.com/mattn/go-isatty"
)

type file struct {
	content  string
	filename string
}

// func (f *file) detect() error {
// 	if lexer != nil {
// 		return nil
// 	}
// 	if f.filename != "-" {
// 		f.lexer = lexers.Match(f.filename)
// 	}
// 	// if we can't match by filename, try analyzing content
// 	if f.lexer == nil {
// 		f.lexer = lexers.Analyse(f.content)
// 	}
// 	// even if that doesn't work, switch to fallback
// 	if f.lexer == nil {
// 		f.lexer = lexers.Fallback
// 	}
// 	f.lexer = chroma.Coalesce(f.lexer)
// 	return nil
// }

func (f *file) format() error {
	var lexer chroma.Lexer

	// detect lexer
	if *lf != "" {
		lexer = lexers.Get(*sf)
		if lexer == nil {
			return fmt.Errorf("unsupported language %s", *sf)
		}
	} else {
		lexer = lexers.Match(f.filename)
		if lexer == nil {
			lexer = lexers.Analyse(f.content)
		}
		if lexer == nil {
			lexer = lexers.Fallback
		}
	}
	lexer = chroma.Coalesce(lexer)
	// detect style

	style := styles.Get(*sf)
	if style == nil {
		return fmt.Errorf("unknown style value %s", *sf)
	}

	iter, err := lexer.Tokenise(nil, f.content)
	if err != nil {
		return err
	}

	log.Info("Using config", "lexer", lexer.Config().Name, "style", style.Name)

	formatter := formatters.Get("terminal16m")
	return formatter.Format(os.Stdout, style, iter)
}

func isAtty() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}
