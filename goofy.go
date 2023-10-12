package goofy

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"runtime"
)

type Goofer struct {
	InfoLog  *log.Logger
	DebugLog *log.Logger
	ErrLog   *log.Logger
}

func NewGoofer() *Goofer {
	return &Goofer{
		InfoLog:  InfoLog,
		DebugLog: DebugLog,
		ErrLog:   ErrLog,
	}
}

var BaseType404Err = errors.New("identifier base type not found")

func getType(ident *ast.Ident) (string, error) {
	for ident.Obj != nil {
		t, ok := ident.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return "", BaseType404Err
		}
		ident, ok = t.Type.(*ast.Ident)
		if !ok {
			return "", BaseType404Err
		}
	}

	return ident.Name, nil
}

func (g *Goofer) Run(src_roots []string, ignore_rgx *regexp.Regexp) error {
	fset := token.NewFileSet()

	for _, root := range src_roots {

		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				return nil
			}

			if ignore_rgx.MatchString(path) {
				fmt.Fprint(g.DebugLog.Writer(), "\x1b[31m")
				g.DebugLog.Printf("%-6s: %v\n", "skip", path)
				fmt.Fprint(g.DebugLog.Writer(), "\x1b[0m")
				return filepath.SkipDir
			}

			fmt.Fprint(g.DebugLog.Writer(), "\x1b[90m")
			g.DebugLog.Printf("%-6s: %v\n", "walk", path)
			fmt.Fprint(g.DebugLog.Writer(), "\x1b[0m")

			pkgs, err := parser.ParseDir(fset, path, nil, 0)
			if err != nil {
				g.ErrLog.Println(err)
				return err
			}

			// typeDecl := map[string]string{}
			for _, pkg := range pkgs {
				fmt.Fprint(g.DebugLog.Writer(), "\x1b[96m")
				g.DebugLog.Printf("\x1b[96m%-6s: %v\n", "pkg", pkg.Name)
				fmt.Fprint(g.DebugLog.Writer(), "\x1b[0m")

				for _, file := range pkg.Files {
					for _, obj := range file.Scope.Objects {
						typespec, ok := obj.Decl.(*ast.TypeSpec)
						if !ok {
							continue
						}

						if typespec.Name.Name == "Timeline" {
							runtime.Breakpoint()
						}
						switch t := typespec.Type.(type) {
						case *ast.StructType:
							if t.Incomplete {
								continue
							}
							g.DebugLog.Printf("struct: %v\n", obj.Name)
							var typ string
							for _, field := range t.Fields.List {
								switch ftype := field.Type.(type) {
								case *ast.Ident:
									typ, err = getType(ftype)
									if err != nil {
										return err
									}
								case *ast.SelectorExpr:
									typ = "selector"
								}
								g.DebugLog.Printf("\t%-10v\ttype: %v\n", field.Names[0].Name, typ)
							}
						}
					}
				}
			}

			return nil
		})

		if err != nil {
			return err
		}

	}
	return nil
}
