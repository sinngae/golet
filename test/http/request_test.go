package http

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo&q=bar&both=x&prio=1&orphan=nope&empty=not",
		strings.NewReader("z=post&both=y&prio=2&=nokey&orphan;empty=&"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	q := req.FormValue("q")
	assert.Equal(t, "foo", q)

	z := req.FormValue("z")
	assert.Equal(t, "post", z)

	bq, found := req.PostForm["q"]
	assert.Equal(t, true, found)
	assert.Equal(t, "", bq)

	bz := req.PostFormValue("z")
	assert.Equal(t, "post", bz)

	if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
		t.Errorf(`req.Form["q"] = %q, want ["foo", "bar"]`, qs)
	}
	if both := req.Form["both"]; !reflect.DeepEqual(both, []string{"y", "x"}) {
		t.Errorf(`req.Form["both"] = %q, want ["y", "x"]`, both)
	}
	if prio := req.FormValue("prio"); prio != "2" {
		t.Errorf(`req.FormValue("prio") = %q, want "2" (from body)`, prio)
	}
	if orphan := req.Form["orphan"]; !reflect.DeepEqual(orphan, []string{"", "nope"}) {
		t.Errorf(`req.FormValue("orphan") = %q, want "" (from body)`, orphan)
	}
	if empty := req.Form["empty"]; !reflect.DeepEqual(empty, []string{"", "not"}) {
		t.Errorf(`req.FormValue("empty") = %q, want "" (from body)`, empty)
	}
	if nokey := req.Form[""]; !reflect.DeepEqual(nokey, []string{"nokey"}) {
		t.Errorf(`req.FormValue("nokey") = %q, want "nokey" (from body)`, nokey)
	}
}
