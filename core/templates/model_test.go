package templates

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"github.com/shiqiyue/go-admin-gen/util"
	"testing"
)

func TestGenEditReq(t *testing.T) {
	templateData := make(map[string]interface{}, 0)
	templateData["PACKAGE"] = "dto"
	d := &dto.Model{
		Name:        "AccountEditReq",
		Description: "联系人修改-参数",
		Fields:      make([]*dto.ModelField, 0),
	}
	d.Fields = append(d.Fields, &dto.ModelField{
		Name:        "Name",
		Description: "名称",
		Type:        "int32",
		Ptr:         false,
	})
	d.Fields = append(d.Fields, &dto.ModelField{
		Name:        "CreatedAt",
		Description: "创建时间",
		Type:        "time.Time",
		Ptr:         true,
	})

	templateData["MODEL"] = d
	r, err := util.DoTemplate(MODEL, "test.go", templateData)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(r))
	}

}
