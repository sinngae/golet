package httplib

import (
	"context"
	"fmt"
	"testing"

	"github.com/sinngae/gland/pkg/httplib/invoke"
	"github.com/stretchr/testify/assert"
)

func TestHttpLib(t *testing.T) {
	t.Run("test httplib", func(t *testing.T) {
		httpCli := invoke.NewHttpClient()
		resp, err := httpCli.Get(context.TODO(), "https://www.baidu.com", nil)
		assert.NoError(t, err)
		fmt.Println(resp)
	})
}

func BenchmarkHttpLib(b *testing.B) {

}

func ExampleHttpLib() {

}
