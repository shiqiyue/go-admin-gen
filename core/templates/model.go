package templates

import (
	"fmt"
	"strings"
)

type Model struct {
	Name string

	ShortName string

	Description string

	Remarks []string

	Fields []*ModelField

	Methods []*ModelMethod
}

type ModelField struct {
	Name string

	Description string

	Type string

	Ptr bool

	Tag string
}

type ModelMethod struct {
	Name        string
	Description string
	Body        string
	Args        []*ModelMethodArg
	Results     []*ModelMethodResult
}

type ModelMethodArg struct {
	Name string
	Type string
	Ptr  bool
}

type ModelMethodResult struct {
	Name string
	Type string
	Ptr  bool
}

// 输出参数
func (m *ModelMethod) OutputArgs() string {
	argStrs := make([]string, 0)
	for _, arg := range m.Args {
		if arg.Ptr {
			argStrs = append(argStrs, fmt.Sprintf("%s *%s", arg.Name, arg.Type))
		} else {
			argStrs = append(argStrs, fmt.Sprintf("%s %s", arg.Name, arg.Type))
		}
	}
	return strings.Join(argStrs, ", ")
}

func (m *ModelMethod) OutputResults() string {
	if len(m.Results) == 0 {
		return ""
	}
	resultStrs := make([]string, 0)
	for _, result := range m.Results {
		if result.Ptr {
			resultStrs = append(resultStrs, fmt.Sprintf("%s *%s", result.Name, result.Type))
		} else {
			resultStrs = append(resultStrs, fmt.Sprintf("%s %s", result.Name, result.Type))
		}
	}
	return fmt.Sprintf("(%s)", strings.Join(resultStrs, ", "))

}
