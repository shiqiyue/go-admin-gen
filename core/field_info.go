package core

import (
	"github.com/iancoleman/strcase"
	"strings"
)

var (
	CREATED_AT_FIELD_NAME string = strings.ToLower("CreatedAt")
	UPDATED_AT_FIELD_NAME string = strings.ToLower("UpdatedAt")
	DELETED_AT_FIELD_NAME string = strings.ToLower("DeletedAt")
	ID_FIELD_NAME         string = strings.ToLower("Id")
)

type FieldInfo struct {
	// 名称
	Name string
	// 类型
	Type string
	// 是否可以为空
	Nullable bool
	// tag
	Tag string
	// Gorm Tag
	GormTag string
}

// gql 字段名称
func (i FieldInfo) GqlFieldName() string {
	return strcase.ToLowerCamel(i.Name)
}

// golang 字段名称
func (i FieldInfo) GoFieldName() string {
	return strcase.ToCamel(i.Name)
}

// golang 字段类型
func (i FieldInfo) GoFieldType() string {
	return i.Type
}

// golang 字段-是否是指针
func (i FieldInfo) GoFieldPtr() bool {
	return i.Nullable
}

// 数据库字段名称
func (i FieldInfo) DBFieldName() string {
	return strcase.ToSnake(i.Name)
}

func (i FieldInfo) IsDetail() bool {
	return strings.Contains(i.Tag, "detail")
}

func (i FieldInfo) IsAdd() bool {
	name := strings.ToLower(i.Name)
	if name == CREATED_AT_FIELD_NAME || name == UPDATED_AT_FIELD_NAME || name == DELETED_AT_FIELD_NAME || name == ID_FIELD_NAME {
		return false
	}

	return true
}

func (i FieldInfo) IsEdit() bool {
	name := strings.ToLower(i.Name)
	if name == CREATED_AT_FIELD_NAME || name == UPDATED_AT_FIELD_NAME || name == DELETED_AT_FIELD_NAME {
		return false
	}
	return true
}

func (i FieldInfo) IsFilter() bool {
	name := strings.ToLower(i.Name)
	if name == ID_FIELD_NAME || name == DELETED_AT_FIELD_NAME {
		return false
	}
	if i.IsArray() {
		return false
	}
	if i.IsJson() {
		return false
	}
	if i.IsTime() {
		return false
	}
	return true
}

func (i FieldInfo) IsVueQuery() bool {
	name := strings.ToLower(i.Name)
	if name == DELETED_AT_FIELD_NAME {
		return false
	}

	return true
}

func (i FieldInfo) IsTableColumn() bool {
	name := strings.ToLower(i.Name)
	if name == ID_FIELD_NAME || name == DELETED_AT_FIELD_NAME {
		return false
	}
	if i.IsJson() {
		return false
	}
	if i.IsArray() {
		return false
	}
	return true
}

func (i FieldInfo) IsSortKey() bool {
	name := strings.ToLower(i.Name)
	if name == DELETED_AT_FIELD_NAME {
		return false
	}
	if i.IsJson() {
		return false
	}
	if i.IsArray() {
		return false
	}
	return true
}

func (i FieldInfo) Scalar() string {
	return GetScalarByType(i.Type)
}

func (i FieldInfo) IsArray() bool {
	if strings.Contains(i.Type, "Array") {
		return true
	}
	return false
}

func (i FieldInfo) IsJson() bool {
	if strings.Contains(i.Type, "JSON") {
		return true
	}
	return false
}

func (i FieldInfo) IsTime() bool {
	return i.Type == "time.Time"
}

func (i FieldInfo) Description() string {
	commentIndex := strings.LastIndex(i.GormTag, "comment:")
	if commentIndex >= 0 {
		endIndex := len(i.GormTag) - 1
		for j := endIndex; j > commentIndex; j-- {
			if i.GormTag[j] == ';' {
				endIndex = j
				break
			}
		}
		return i.GormTag[commentIndex+8 : endIndex+1]
	}
	return ""
}
