package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"text/template"
)

func WriteTemplate(t, filename string, data interface{}) error {
	bs, err := DoTemplate(t, filename, data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bs, 0777)
	if err != nil {
		return err
	}
	if path.Ext(filename) == ".go" {
		return RunInteractive(fmt.Sprintf("goimports -w %s", filename))
	}
	return nil
}

func DoTemplate(t, filename string, data interface{}) ([]byte, error) {
	temp, err := template.New(filename).Parse(t)
	if err != nil {
		return nil, err
	}
	var content bytes.Buffer
	writer := io.Writer(&content)

	err = temp.Execute(writer, &data)
	if err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}
