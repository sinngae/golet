package stdst

import "log"

type (
	BSt struct {
		v0 int
	}
	DSt struct {
		v0 string
		BSt
	}
)

func (obj *BSt) Fun() int {
	return 0
}

func (obj *DSt) Fun() int {
	log.Default()
	return 1
}

func ExampleClass() {
	val := DSt{}
	v := val.Fun()
	v1 := val.v0
	println(v, v1)
	//output:
}
