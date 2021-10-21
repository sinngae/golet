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

func ExampleXml() {
	type OrderNo struct {
		OrderNo string `xml:"OrderNo"`
	}
	type Data struct {
		XmlName  xml.Name  `xml:"Doc"`
		ParentId string    `xml:"ParentId"`
		EshopId  string    `xml:"EshopId"`
		OrderNos []OrderNo `xml:"OrderNos"`
	}
	type Doc struct {
		Data Data `xml:"Doc"`
	}
	doc := Doc{
		Data{
			ParentId: "test",
			EshopId:  "eshop",
			OrderNos: []OrderNo{{"123"}, {"2312"}},
		},
	}
	got, err := xml.MarshalIndent(doc, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "")
	fmt.Println(got)
	fmt.Println(err)
	//output:

}
