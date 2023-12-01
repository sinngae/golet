package zip

import (
	"archive/zip"
	"log"
	"os"
)

type WriteTest struct {
	Name   string
	Data   []byte
	Method uint16
	Mode   os.FileMode
}

var writeTests = []WriteTest{
	{
		Name:   "foo",
		Data:   []byte("Rabbits, guinea pigs, gophers, marsupial rats, and quolls."),
		Method: zip.Store,
		Mode:   0666,
	},
	{
		Name:   "bar",
		Data:   nil, // large data set in the test
		Method: zip.Deflate,
		Mode:   0644,
	},
	{
		Name:   "setuid",
		Data:   []byte("setuid file"),
		Method: zip.Deflate,
		Mode:   0755 | os.ModeSetuid,
	},
	{
		Name:   "setgid",
		Data:   []byte("setgid file"),
		Method: zip.Deflate,
		Mode:   0755 | os.ModeSetgid,
	},
	{
		Name:   "symlink",
		Data:   []byte("../link/target"),
		Method: zip.Deflate,
		Mode:   0755 | os.ModeSymlink,
	},
}

func ExampleZipFile() {
	file, err := os.Create("./test.zip")
	if err != nil {
		log.Fatal(err)
	}
	//buf := new(file)
	w := zip.NewWriter(file)

	for _, wt := range writeTests {
		header := &zip.FileHeader{
			Name:   wt.Name,
			Method: wt.Method,
		}
		if wt.Mode != 0 {
			header.SetMode(wt.Mode)
		}
		f, err := w.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(wt.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	println("end")

	// output:

}
