package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

/*
go test -run ''      # 运行所有测试
go test -run Foo     # 匹配 Foo 相关的顶级测试，如 TestFooBar
go test -run Foo/A=  # 匹配 Foo 相关的顶级测试, 并匹配子测试 A=
go test -run /A=1    # 匹配所有顶级测试, 并匹配它们的子测试 A=1
*/
var testcases = map[string]map[string][]string{
	"test": {
		"case1": []string{"test-case1-a", "test-case1-b", "test-case1-c"},
		"case2": []string{"test-case2-d", "test-case2-e", "test-case2-f"},
	},
	"live": {
		"case1": []string{"live-case1-a", "live-case1-b", "live-case1-c"},
		"case2": []string{"live-case2-d", "live-case2-e", "live-case2-f"},
	},
}

func TestDemo(t *testing.T) {
	t.Run("A=1 abc", func(t *testing.T) { // 子测试，同级子测试是并行的
		err := fmt.Errorf("this a test A")
		fmt.Println(err)
	})
	t.Run("B=2,C=1", func(t *testing.T) {
		err := fmt.Errorf("this a test1")
		fmt.Println(err)
	})
	t.Run("B=3,C=2", func(t *testing.T) {
		err := fmt.Errorf("this a test2")
		fmt.Println(err)
	})
	for env, tcm := range testcases {
		for cond, tcs := range tcm {
			for idx, tc := range tcs {
				t.Run(fmt.Sprintf("env=%s,case=%s,idx=%d", env, cond, idx), func(t *testing.T) {
					fmt.Printf("%s\n", tc)
				})
			}
		}
	}
}

func BenchmarkDemo0(b *testing.B) {
	str := "abcdefg"
	b.ReportAllocs() // -test.benchmem
	//b.ReportMetric() //?
	b.ResetTimer() // reset timer & mem counter
	for i := 0; i < b.N; i++ {
		_ = len(str)
	}
}

// 并行压测，不限CPU
// `go test -cpu`
func BenchmarkDemo(b *testing.B) {
	tmpl := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			err := tmpl.Execute(&buf, "world")
			assert.NoError(b, err)
		}
	})
}

func Example() {
	fmt.Printf("hi, work\n")
}

func TestMain(m *testing.M) {
	os.Exit(m.Run()) // Run call main or testcases?
}