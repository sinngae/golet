package scan

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func ExampleScan() {
	filename := path.Join("data", "demo.txt")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		args := strings.Split(line, "?")
		orderNos := strings.Split(args[1], "\\u0026")
		orderNos = orderNos[:len(orderNos)-2]
	}

	if err := scan.Err(); err != nil {
		panic(err)
	}
}
