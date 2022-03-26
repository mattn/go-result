//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

var tmpl = `
type Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any] struct {
    {{- range $val := .}}
	v{{$val}}  T{{$val}}
    {{- end }}
	err error
}

func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Unwrap() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		panic(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Err() error {
	return r.err
}

func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IsErr() bool {
	return r.err != nil
}

func Ok{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any]({{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}} T{{$val}}{{ end }}) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
    return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{{"{"}}{{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}}: v{{$val}}{{ end }}, err: nil}
}

func Err{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any](err error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{err: err{{"}"}}
}

{{- $ns := .}}
{{ range $val := $ns}}
func (r *Result{{len $ns}}[{{ range $val := $ns}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) V{{$val}}() T{{$val}} {
	return r.v{{$val}}
}
{{ end }}
`

func main() {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "package result\n\n")

	var ns []int
	t := template.Must(template.New("").Parse(tmpl))
	for i := 1; i <= 10; i++ {
		ns = append(ns, i)
		err := t.Execute(&buf, ns)
		if err != nil {
			log.Fatal(err)
		}
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("struct.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(b)
}
