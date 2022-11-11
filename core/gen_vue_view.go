package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"github.com/shiqiyue/go-admin-gen/util"
	"path"
	"strings"
)

func (c *GenContext) genVueView() error {
	v := &templates.ViewVue{
		Filters:       make([]*templates.ViewVueFilter, 0),
		TableColumns:  make([]*templates.ViewTableColumn, 0),
		EditFormItems: make([]*templates.ViewTableEditFormItem, 0),
	}
	queryFields := make([]string, 0)
	for _, field := range c.Fields {
		if field.IsFilter() {
			v.Filters = append(v.Filters, &templates.ViewVueFilter{
				FieldLabel:  field.Name,
				FieldName:   field.GqlFieldName(),
				FieldName2:  "",
				ControlType: templates.CONTROL_TYPE_INPUT,
			})
		}
		if field.IsVueQuery() {
			queryFields = append(queryFields, field.GqlFieldName())
		}

		if field.IsAdd() {
			rules := ""
			if !field.Nullable {
				rules = "[{required: true, message: '不能为空'}]"
			}
			v.EditFormItems = append(v.EditFormItems, &templates.ViewTableEditFormItem{
				FieldLabel:  field.Name,
				FieldName:   field.GqlFieldName(),
				Rules:       rules,
				ControlType: templates.CONTROL_TYPE_INPUT,
			})
		}

		if field.IsTableColumn() {
			filter := ""
			if field.Type == "time.Time" {
				filter = "parseDateTime"
			}
			v.TableColumns = append(v.TableColumns, &templates.ViewTableColumn{
				FieldLabel: field.Name,
				FieldName:  field.GqlFieldName(),
				Filter:     filter,
			})
		}
	}
	v.SearchGqlName = c.graphqlPageQueryName()
	v.SearchGql = fmt.Sprintf(`query %s($data: %s!){
        %s(data: $data) {
            total
            records {
                %s
            }
        }
	}`, v.SearchGqlName, c.graphqlPageInputName(), v.SearchGqlName, strings.Join(queryFields, "\n"))
	v.AddGql = fmt.Sprintf(`mutation %s($data: %s!){
                %s(data: $data)
            }`, c.graphqlAddMutationName(), c.graphqlAddReqName(), c.graphqlAddMutationName())
	v.EditGql = fmt.Sprintf(`mutation %s($data: %s!){
                %s(data: $data)
            }`, c.graphqlEditMutationName(), c.graphqlEditReqName(), c.graphqlEditMutationName())
	v.RemovesSql = fmt.Sprintf(`mutation %s($data: %s!){
                %s(data: $data)
            }`, c.graphqlRemovesMutationName(), c.graphqlRemoveReqName(), c.graphqlRemovesMutationName())
	rs, err := util.DoTemplate(templates.VIEW_VUE, "view.vue", v)
	if err != nil {
		return err
	}
	filePath := path.Join(c.Cfg.GetVueViewDir(), c.ModelSneakName()+".vue")

	err = util.WriteFile(rs, filePath, false)
	if err != nil {
		return err
	}
	return nil
}
