package test

func ExampleDemo() {
	life := struct {
		Cheated func(int) bool
	}{}
	life.Cheated = func(val int) bool {
		return val == 0
	}
	you := 1
Dance:
	for { //keep going
		if life.Cheated(you) {
			for { // will you linger at this
				break // or just have a break
			}
			goto Dance // and goto dance
		}
	}
}
