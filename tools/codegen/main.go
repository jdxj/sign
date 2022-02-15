package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	WriteToFile(RenderCodes(ParseCodes()))
}

type Code struct {
	Name string
	HTTP int
	Desc string
}

func ParseCodes() (codes []Code) {
	var (
		fileSet  = token.NewFileSet()
		filename = os.Getenv("GOFILE")
	)
	astFile, err := parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("ParseFile: %s", err)
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		genDecl, ok := node.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			return true
		}

		// 是否找到该 const 声明中, type 为 'code' 的声明
		find := false
		for _, spec := range genDecl.Specs {
			valueSpec := spec.(*ast.ValueSpec)
			ident, ok := valueSpec.Type.(*ast.Ident)
			if ok {
				if ident.Name == "code" {
					find = true
				} else {
					// 发现新类型时认为未找到
					find = false
				}
			}
			if !find {
				continue
			}

			// 解析 const 上的注释
			// 过滤 '_'
			// len(Names) 必定大于0
			if valueSpec.Names[0].Name == "_" {
				continue
			}

			// 注释是必须的
			if valueSpec.Doc == nil {
				log.Fatalf("const '%s' has no comment\n", valueSpec.Names[0].Name)
			}

			// len(List) 必定大于0, 只解析第一行
			txt := valueSpec.Doc.List[0].Text
			code := parseCode(txt)
			codes = append(codes, code)
		}
		return true
	})
	return
}

func parseCode(txt string) Code {
	txt = strings.TrimPrefix(txt, "//")
	txt = strings.TrimSpace(txt)

	const num = 3
	parts := strings.SplitN(txt, " ", num)
	if len(parts) != num {
		log.Fatalf("format err: %s, correct: '// name http desc'\n", txt)
	}
	http, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("invalid http '%s' in '%s'", parts[1], txt)
	}

	return Code{
		Name: parts[0],
		HTTP: http,
		Desc: parts[2],
	}
}

func RenderCodes(codes []Code) []byte {
	t, err := template.New("").Parse(tplSrc)
	if err != nil {
		log.Fatalf("Parse: %s", err)
	}

	var (
		buf  = bytes.NewBuffer(nil)
		data = map[string]interface{}{
			"pkg":   os.Getenv("GOPACKAGE"),
			"codes": codes,
		}
	)
	err = t.Execute(buf, data)
	if err != nil {
		log.Fatalf("Execute: %s\n", err)
	}

	body, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Source: %s\n", err)
	}
	return body
}

func WriteToFile(data []byte) {
	var (
		goFile   = os.Getenv("GOFILE")
		ext      = filepath.Ext(goFile)
		filename = strings.TrimSuffix(goFile, ext) + "_generated" + ext
	)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Create: %s\n", err)
	}
	defer func() {
		_ = file.Sync()
		_ = file.Close()
	}()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Write: %s\n", err)
	}
}
