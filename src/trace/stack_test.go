package trace

import (
	"fmt"
	"testing"
)

func TestCallers(t *testing.T) {
	pcList := Callers()
	for _, pc := range *pcList {
		file, line, fnName := fileLineFnName(pc)
		//fmt.Println(file, line, funcName(fnName))
		fmt.Println(file, line, fnName)
	}
}

func TestProgramStack_MarshalText(t *testing.T) {
	pcList := Callers()

	str := pcList.MarshalText()
	fmt.Println(string(str))
}

func TestProgramStack_Format(t *testing.T) {
	pcList := Callers()

	fmt.Printf("%+v", pcList)
}
