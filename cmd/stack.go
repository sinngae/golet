package main

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
)

type myst struct {
	abc int
	def string
}

func main() {
	Example(&myst{
		abc: 1,
		def: "test",
	})
}

/*
buf:{{{goroutine 1 [running]:
runtime/debug.Stack(0xc00007eef0, 0x1037037, 0x1167f80)
        /usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/debug/stack.go:24 +0x9d
main.Example0(0xc00007ef38)
        /Users/ziqiangren/gitwork/gland/cmd/stack.go:23 +0x26
main.main()
        /Users/ziqiangren/gitwork/gland/cmd/stack.go:15 +0x5a
}}}

*/
func Example0(st *myst) {
	//debug.PrintStack()
	buf := debug.Stack()
	fmt.Printf("buf:{{{%s}}}", buf)
}

/*
buf:{{{goroutine 1 [running]:
main.Example1(0xc00007ef38)
        /Users/ziqiangren/gitwork/gland/cmd/stack.go:40 +0x6d
main.main()
        /Users/ziqiangren/gitwork/gland/cmd/stack.go:15 +0x5a
}}}

*/
func Example1(st *myst) {
	buf := make([]byte, 2000)
	runtime.Stack(buf, false)
	fmt.Printf("buf:{{{%s}}}\n", buf)

	//runtime.StartTrace()
	//runtime.StopTrace()
}

func Example2(st *myst) {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Printf("caller0:{{{%v,%s,%d,%t}}}\n", pc, file, line, ok)
}

/*
caller0:{{{
/Users/ziqiangren/gitwork/gland/cmd/stack.go:15 (0x109b629)
/usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/proc.go:203 (0x102aec5)
/usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/asm_amd64.s:1357 (0x10533b0)
}}}

*/
func Example(st *myst) {
	stack := Stack(1)
	fmt.Printf("caller0:{{{\n%s}}}\n", stack)
}

func Stack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.Bytes()
}
