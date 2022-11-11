package templates

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/util"
	"testing"
)

func TestGenViewVueView(t *testing.T) {
	v := &ViewVue{
		Filters:       make([]*ViewVueFilter, 0),
		TableColumns:  make([]*ViewTableColumn, 0),
		EditFormItems: make([]*ViewTableEditFormItem, 0),
	}
	v.Filters = append(v.Filters, &ViewVueFilter{
		FieldLabel:  "名称",
		FieldName:   "name",
		ControlType: CONTROL_TYPE_INPUT,
	})

	v.TableColumns = append(v.TableColumns, &ViewTableColumn{
		FieldLabel: "名称",
		FieldName:  "name",
		Filter:     "",
	})
	v.TableColumns = append(v.TableColumns, &ViewTableColumn{
		FieldLabel: "创建时间",
		FieldName:  "createdAt",
		Filter:     "dateFormatter",
	})

	v.EditFormItems = append(v.EditFormItems, &ViewTableEditFormItem{
		FieldLabel:  "名称",
		FieldName:   "name",
		Rules:       "[{required: true, message: '不能为空'}]",
		ControlType: CONTROL_TYPE_INPUT,
	})

	rs, err := util.DoTemplate(VIEW_VUE, "view.vue", v)
	if err != nil {
		t.Error(err)
		return
	}
	rsStr := string(rs)
	fmt.Println(rsStr)
	util.WriteFile(rs, "t.vue", true)
}
