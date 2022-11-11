package templates

var (
	CONTROL_TYPE_INPUT = "input"

	CONTROL_TYPE_DATE_PICKER = "date-picker"

	CONTROL_TYPE_DATE_TIME_PICKER = "date-time-picker"

	CONTROL_TYPE_DATE_RANGE_PICKER = "date-range-picker"
)

type ViewVue struct {
	// 过滤项
	Filters []*ViewVueFilter
	// 表格项
	TableColumns []*ViewTableColumn

	EditFormItems []*ViewTableEditFormItem

	SearchGql string

	SearchGqlName string

	AddGql string

	EditGql string

	RemovesSql string
}

type ViewVueFilter struct {
	// 字段Label
	FieldLabel string
	// 字段名称
	FieldName string
	// 字段名称
	FieldName2 string
	// 控件类型
	ControlType string
}

type ViewTableColumn struct {
	// 字段Label
	FieldLabel string
	// 字段名称
	FieldName string
	// 过滤器
	Filter string
}

type ViewTableEditFormItem struct {
	// 字段Label
	FieldLabel string
	// 字段名称
	FieldName string
	// 规则
	Rules string
	// 控件类型
	ControlType string
}
