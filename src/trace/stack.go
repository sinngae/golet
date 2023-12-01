package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// ProgramStack is a stack of Program Counters
//	PC, as known as PC register of CPU, point to function address
type ProgramStack []uintptr

func Callers() *ProgramStack {
	const depth = 32 // FIXME max depth, extra stack will be ignored
	var programCountList [depth]uintptr
	n := runtime.Callers(3, programCountList[:])
	var st ProgramStack = programCountList[0:n]
	return &st
}

func (ps *ProgramStack) MarshalText() []byte {
	buf := new(bytes.Buffer)
	for _, pc := range *ps {
		file, line, fnName := fileLineFnName(pc)
		fmt.Fprintf(buf, "%s %s:%d", funcName(fnName), file, line)
	}
	return buf.Bytes()
}

//func (ps *ProgramStack) Format(st fmt.State, verb rune) {
//	switch verb {
//	case 'v':
//		switch {
//		case st.Flag('+'):
//			txt, _ := ps.MarshalText()
//			fmt.Fprint(st, txt)
//		}
//	}
//}

func fileLineFnName(pc uintptr) (string, int, string) {
	_pc := pc - 1
	fn := runtime.FuncForPC(_pc)
	if fn == nil {
		return "unknown", 0, "unknown"
	}
	file, line := fn.FileLine(_pc)
	return file, line, fn.Name()
}

func funcName(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
