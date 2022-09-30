package stdst

import "runtime"

func ExampleChannel() {
	ch := make(chan struct{})
	close(ch)
	runtime.Breakpoint()
	//output:
}
