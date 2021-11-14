package xml

import (
	"encoding/xml"
	"fmt"
)

func ExampleMarshalIndent() {
	type Data struct {
		ParentId string `xml:"ParentId"`
		EshopId  string `xml:"EshopId"`
	}
	type Doc struct {
		Data Data `xml:"Doc"`
	}
	doc := Doc{
		Data{
			ParentId: "test",
			EshopId:  "eshop",
		},
	}
	got, err := xml.MarshalIndent(doc, "", "")
	fmt.Println(got)
	fmt.Println(err)
	//output:

}

func ExampleXmlMarshal() {
	const header = `<?xml version="1.0" encoding="UTF-8"?>`
	type OrderNo struct {
		OrderNo string `xml:"OrderNo"`
	}
	type Doc struct {
		ParentId string    `xml:"ParentId"`
		EshopId  string    `xml:"EshopId"`
		OrderNos []OrderNo `xml:"OrderNos"`
	}
	//type Doc struct {
	//	Data Data `xml:"Doc"`
	//}
	doc := Doc{
		ParentId: "test",
		EshopId:  "eshop",
		OrderNos: []OrderNo{{"123"}, {"2312"}},
	}
	got, err := xml.MarshalIndent(doc, "", "")
	fmt.Println(header + string(got))
	fmt.Println(err)
	//output:
}

func ExampleXmlUnmarshal() {
	const header = `<?xml version="1.0" encoding="UTF-8"?>`
	type OrderNo struct {
		OrderNo string `xml:"OrderNo"`
	}
	type Doc struct {
		ParentId string    `xml:"ParentId"`
		EshopId  string    `xml:"EshopId"`
		OrderNos []OrderNo `xml:"OrderNos"`
	}
	//type Doc struct {
	//	Data Data `xml:"Doc"`
	//}
	doc := Doc{
		ParentId: "test",
		EshopId:  "eshop",
		OrderNos: []OrderNo{{"123"}, {"2312"}},
	}
	got, err := xml.Unmarshal(doc, "", "")
	fmt.Println(header + string(got))
	fmt.Println(err)
	//output:

}
