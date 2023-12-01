package template_test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	testcases := []struct {
		tmpl string
		data interface{}
		want string
	}{
		{`{{.Val}}`, struct{ Val string }{"abc"}, "abc"},
		{`{{if not .Val}}0{{else}}1{{end}}`, struct{ Val bool }{true}, "1"},
		{`{{if not .Val}}0{{else}}1{{end}}`, struct{ Val *int32 }{nil}, "0"},
		//{`{{if .Val 0 else 1 end}}`, struct{ Val *int32 }{nil}, "1"},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			//tmpl := template.Must(template.New("json_demo").Parse(jsonTpl))
			tmpl, err := template.New("json_demo").Parse(tc.tmpl)
			assert.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tc.data)
			assert.NoError(t, err)
			str := buf.String()
			assert.Equal(t, tc.want, str)
		})
	}
}
