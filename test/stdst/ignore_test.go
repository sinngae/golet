package stdst

import "fmt"

func demo() (int, error) {
	return 0, fmt.Errorf("test")
}

func ExampleDemo() {
	a, _ := demo()
	fmt.Println(a)
	// output:
	//0
}
