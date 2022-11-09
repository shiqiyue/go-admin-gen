package templates

var MODEL = `
package {{.PACKAGE}}
import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// {{.MODEL.Name}} {{.MODEL.Description}}
type {{.MODEL.Name}} struct {
{{- range .MODEL.Fields }}
	// {{.Name}} {{.Description}}
	{{if .Ptr}}{{.Name}} *{{.Type}} {{else}}{{.Name}} {{.Type}} {{end}}
{{ end }}
}
`
