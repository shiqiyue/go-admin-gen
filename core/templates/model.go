package templates

var MODEL = `
package {{.PACKAGE}}
import (
{{- range .INPUTS }}
	"{{.}}"
{{end}}
)

{{ range $MODEL := .MODELS }}
// {{$MODEL.Name}} {{$MODEL.Description}}
type {{$MODEL.Name}} struct {
{{range $Field:= $MODEL.Fields }}
	// {{$Field.Name}} {{$Field.Description}}
	{{if $Field.Ptr}}{{$Field.Name}} *{{$Field.Type}} {{else}}{{$Field.Name}} {{$Field.Type}} {{end}} {{$Field.Tag}}
{{ end }}
}

{{range $METHOD:= $MODEL.Methods }}
// {{$METHOD.Name}} {{$METHOD.Description}}
func ({{$MODEL.ShortName}} *{{$MODEL.Name}}) {{$METHOD.Name}}({{$METHOD.OutputArgs}}){{$METHOD.OutputResults}} {
	{{$METHOD.Body}}
}
{{end}}
{{ end }}

`
