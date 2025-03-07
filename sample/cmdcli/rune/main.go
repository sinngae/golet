package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := ioutil.ReadFile("demo_th.xlsx")
	if err != nil {
		log.Fatalf("exit 0, err=%v", err)
		return
	}
	xlsFile, err := excelize.OpenReader(bytes.NewBuffer(file))
	if err != nil {
		log.Fatalf("exit 1, err=%v", err)
		return
	}

	src, err := xlsFile.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("exit 3, err=%v", err)
		return
	}
	sloTnList := make([]string, 0)
	for _, item := range src[1:] {
		if hasFourByteRune(item[2]) {
			sloTnList = append(sloTnList, item[0])
		}
	}

	dest := strings.Join(sloTnList, ",")
	if err := ioutil.WriteFile("out.txt", []byte(dest), 0644); err != nil {
		log.Fatalf("exit 2, err=%v", err)
		return
	}
}

func hasFourByteRune(src string) bool {
	rnList := []rune(src)
	for _, item := range rnList {
		str := string(item)
		btList := []byte(str)
		//fmt.Printf(strconv.Itoa(len(btList)))
		if len(btList) >= 4 {
			return true
		}
	}
	return false
}
