package stdst

type BaseSt struct {
	id  int
	val string
}

type St struct {
	*BaseSt
	flag bool
}

func ExampleObj() {
	st := &St{BaseSt: &BaseSt{id: 1, val: "hi"}}
	var bst *BaseSt
	bst = st.BaseSt

	println(bst)
	// output:

}
