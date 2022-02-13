package main

const (
	tplSrc = `// Code generated by "sign/tools/code"; DO NOT EDIT.

package {{.pkg}}

func init() {
	{{- range .codes}}
	register({{.HTTP}}, {{.Name}}, "{{.Desc}}")
	{{- end}}
}
`
)
