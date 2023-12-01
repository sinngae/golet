package main

import (
	"fmt"
	"strconv"
)

func ExampleRune() {
	data := `🤦🏻‍♂️สุดคุ้ม❗️ลด 50 บาท✨💙 พร้อมส่ง เซ็ตผ้าปูที่นอน พร้อมผ้านวม สีฟ้า นุ่มนิ่ม,5 ฟุต`
	//data := `Huỳnh Khương An Phường 5, Phường 5, Quận Gò Vấp, TP. Hồ Chí Minh`
	rnList := []rune(data)
	for _, item := range rnList {
		str := string(item)
		btList := []byte(str)
		fmt.Printf(strconv.Itoa(len(btList)))
		if len(btList) > 4 {
			println("")
		}
	}
	println(`this is end`)
	// output:

}

func ExampleRune2() {
	data := `สินสยามวิลเลจ 29/44 หมู่3 ตำบล กร่ำ`
	//data := `🤦🏻‍♂️สุดคุ้ม❗️ลด 50 บาท✨💙 พร้อมส่ง เซ็ตผ้าปูที่นอน พร้อมผ้านวม สีฟ้า นุ่มนิ่ม,5 ฟุต`
	//data := `[{"sku": "", "brand": "", "width": 15, "height": 5, "length": 15, "weight": 300, "dg_type": 1, "item_id": 15020563458, "category": "Home & Living", "model_id": 150958663929, "order_sn": "220324Q71THES8", "quantity": 2, "tax_info": [{}], "item_name": "ไฟประดับ สายไฟ ไฟประดับใส่ถ่าน LED ขดลวด กันน้ำ5เมตร ไฟตกแต่ง ไฟLED เลือกสีได้ มี3จังหวะ ปิดไฟ ไฟกระพริบ และไฟสว่างนิ่ง", "shop_name": "C Advance", "model_name": "สีวอม,5เมตร", "net_amount": 0, "order_item": {"item_snapshot": {"item": {"sku": "", "name": "ไฟประดับ สายไฟ ไฟประดับใส่ถ่าน LED ขดลวด กันน้ำ5เมตร ไฟตกแต่ง ไฟLED เลือกสีได้ มี3จังหวะ ปิดไฟ ไฟกระพริบ และไฟสว่างนิ่ง", "country": "TH", "ext_info": {"logistics_info": "{\"7000\":{\"cover_shipping_fee\":false,\"enabled\":true}}"}, "category_id": 57}, "model": {"sku": "", "name": "สีวอม,5เมตร"}}}, "category_id": 18399, "weight_unit": "g", "order_amount": 176, "product_desc": "", "product_name": "ไฟประดับ สายไฟ ไฟประดับใส่ถ่าน LED ขดลวด กันน้ำ5เมตร ไฟตกแต่ง ไฟLED เลือกสีได้ มี3จังหวะ ปิดไฟ ไฟกระพริบ และไฟสว่างนิ่ง", "sub_category": "Lighting", "declare_value": 88, "order_quantity": 2, "declare_name_cn": "Home & Living-Lighting", "declare_name_en": "Home & Living-Lighting", "category_display": "เครื่องใช้ในบ้าน", "current_category": "Home & Living", "sub_sub_category": "Smart Lighting", "declare_value_usd": 2.71, "single_item_price": 88, "declare_name_local": "เครื่องใช้ในบ้าน-โคมไฟและอุปกรณ์ให้แสงสว่าง", "global_category_L1": "Home & Living", "global_category_L2": "Lighting", "global_category_L3": "", "global_category_L4": "", "global_category_L5": "", "global_category_id": 100719, "seller_listing_photo": "", "sub_category_display": "โคมไฟและอุปกรณ์ให้แสงสว่าง", "global_category_id_L1": 100636, "global_category_id_L2": 100719, "global_category_id_L3": 0, "global_category_id_L4": 0, "global_category_id_L5": 0, "global_declare_name_en": "Home & Living-Lighting", "current_category_display": "เครื่องใช้ในบ้าน", "sub_sub_category_display": "ไฟอัจฉริยะ ไฟไร้สาย", "global_declare_name_local": "เครื่องใช้ในบ้าน-โคมไฟและอุปกรณ์ให้แสงสว่าง", "global_category_display_L1": "เครื่องใช้ในบ้าน", "global_category_display_L2": "โคมไฟและอุปกรณ์ให้แสงสว่าง", "global_category_display_L3": "", "global_category_display_L4": "", "global_category_display_L5": ""}]`
	//data := `Huỳnh Khương An Phường 5, Phường 5, Quận Gò Vấp, TP. Hồ Chí Minh`
	rnList := []rune(data)
	rnListDest := make([]rune, 0)
	for _, item := range rnList {
		str := string(item)
		fmt.Printf("%s,%d,%x,%U\n", str, item, item, item)
		if item >= 57344 && item <= 63743 {
			rnListDest = append(rnListDest, item)
		}
	}
	println(`this is end`)
	// output:
}
