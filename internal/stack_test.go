package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"
)
func Stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		//fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
		fmt.Fprintf(buf, "\t%v: %s, %v\n", pc, lines, line)
	}
	return buf.Bytes()
}


/*
buf 10*1024
buf: goroutine 6 [running]:
github.com/sinngae/goerrcode/internal.TestGetStack(0xc000126120)
	/Users/ziqiangren/gitwork/goerrcode/internal/stack_test.go:41 +0x79
testing.tRunner(0xc000126120, 0x114b2e8)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1039 +0xdc
created by testing.(*T).Run
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1090 +0x372

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000126120, 0x114313a, 0xc, 0x114b2e8, 0x1075ee6)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1091 +0x399
testing.runTests.func1(0xc000126000)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1334 +0x78
testing.tRunner(0xc000126000, 0xc00011de10)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1039 +0xdc
testing.runTests(0xc00000c080, 0x1239260, 0x1, 0x1, 0x0)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1332 +0x2a7
testing.(*M).Run(0xc000114000, 0x0)
	/usr/local/Cellar/go/1.14.7/libexec/src/testing/testing.go:1249 +0x1b7
main.main()
	_testmain.go:44 +0x135

buf 100
buf: goroutine 6 [running]:
github.com/sinngae/goerrcode/internal.TestGetStack(0xc000126120)
	/Users/ziqi
*/
func TestGetStack(t *testing.T) {
	testList := []string{"test", "test1"}
	for _, str := range testList {
		myFunc(str)
	}
}

func myFunc(test string) {
	buf := make([]byte, 2000)
	runtime.Stack(buf, true)
	fmt.Printf("buf %s: %s\n", test, buf)
}
