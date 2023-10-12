package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ar-tty/goofy"
)

type patterns []string

func (p *patterns) String() string {
	return fmt.Sprint(*p)
}

func (p *patterns) Set(value string) error {
	*p = append(*p, value)
	return nil
}

func main() {
	g := goofy.NewGoofer()

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "usage: goofy [<options>] path\n\n")
		flag.PrintDefaults()
	}
	flag.BoolFunc("v", "verbose\n", func(s string) error {
		g.DebugLog.SetOutput(os.Stdout)
		return nil
	})

	var ignore_dirs patterns
	flag.Var(&ignore_dirs, "I", "ignore directory\ne.g. goofy -I ./data/ -I ./dbbackup/\n")
	flag.Parse()
	src_roots := flag.Args()

	if len(src_roots) == 0 {
		flag.Usage()
		fmt.Fprint(g.ErrLog.Writer(), "\x1b[31m")
		g.ErrLog.Fatalf("source not set!\n")
		fmt.Fprint(g.ErrLog.Writer(), "\x1b[0m")
	}

	g.DebugLog.Printf("%-15s: %v\n", "ignore dirs", ignore_dirs)
	g.DebugLog.Printf("%-15s: %v\n", "filepaths", src_roots)

	var ptrn string
	for _, p := range ignore_dirs {
		p = regexp.QuoteMeta(filepath.Clean(p))
		ptrn = fmt.Sprintf("%v%v|", ptrn, p)
	}
	ptrn = ptrn[:len(ptrn)-1]

	ignore_rgx, err := regexp.Compile(ptrn)
	if err != nil {
		g.ErrLog.Fatalln(err)
	}
	g.DebugLog.Printf("%-15s: %v\n", "ignore regx", ignore_rgx)
	g.Run(src_roots, ignore_rgx)
}
