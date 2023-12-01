package xls

import (
	"encoding/json"
	"log"
)

func ExampleJson2Xls() {
	jsonStr := `[    {        "name": "Channel Management / CDT Setting",        "permissionCode": "channel_management.cdt_setting.read"    },    {        "name": "Channel Management / LM CDT Setting",        "permissionCode": "channel_management.lm_cdt_setting.read"    },    {        "name": "Channel Management / Campaign Degrade for CDT",        "permissionCode": "channel_management.cdt_of_location.read"    },    {        "name": "Channel Management / Non-working Days",        "permissionCode": "channel_management.non-working_day.read"    },    {        "name": "Channel Management / Channel Flag",        "permissionCode": "channel_management.channel_flag.read"    },    {        "name": "Channel Management / Serviceable Area",        "permissionCode": "channel_management.serviceable_area.read"    },    {        "name": "Channel Management / Channel Split",        "permissionCode": "channel_management.channel_split.read"    },    {        "name": "Channel Management / Shipping Fee Rules",        "permissionCode": "channel_management.shipping_fee_rule.read"    },    {        "name": "Channel Management / Pick Up Configuration",        "permissionCode": "channel_management.pick_up_configuration.read"    },    {        "name": "Channel Management / Branch List",        "permissionCode": "channel_management.branch_list.read"    },    {        "name": "Channel Management / Service Type",        "permissionCode": "channel_management.service_type.read"    },    {        "name": "Channel Management / Channel Config(Integrated)",        "permissionCode": "channel_management.channel_config_integrated.read"    },    {        "name": "Channel Management / Channel Config(Non-integrated)",        "permissionCode": "channel_management.channel_config_non-integrated.read"    },    {        "name": "Channel Management / Route - Service Config",        "permissionCode": "channel_management.service_type_config.read"    },    {        "name": "Channel Management / DG/Non-DG Config",        "permissionCode": "channel_management.dg_config.read"    },    {        "name": "Masked Channel Management / Routing channel",        "permissionCode": "mask_channel_management.routing_channel.read"    },    {        "name": "Masked Channel Management / Lane Service Code",        "permissionCode": "mask_channel_management.lane_service_code.read"    },    {        "name": "Masked Channel Management / Soft Criteria",        "permissionCode": "mask_channel_management.soft_criteria.read"    },    {        "name": "Masked Channel Management / Local Routing Rule",        "permissionCode": "mask_channel_management.routing_rules.read"    },    {        "name": "Masked Channel Management / Masked Channel Configuration",        "permissionCode": "mask_channel_management.configuration.read"    },    {        "name": "Masked Channel Management / Channel Priority & Toggle",        "permissionCode": "mask_channel_management.channel_priority_toggle.read"    },    {        "name": "Masked Channel Management / Local Soft Criteria",        "permissionCode": "mask_channel_management.local_soft_criteria.read"    },    {        "name": "Masked Channel Management / Service Code",        "permissionCode": "mask_channel_management.service_code.read"    },    {        "name": "Masked Channel Management / Receiver Zone",        "permissionCode": "routing_forecasting.forecasting_task.read"    },    {        "name": "Masked Channel Management / Local Volume Routing",        "permissionCode": "mask_channel_management.local_volume_routing.read"    },    {        "name": "Admin Logistics / Delivery Leadtime Settings",        "permissionCode": "admin_logistics.delivery_leadtime_setting.read"    },    {        "name": "Admin Logistics / Whitelist / Blacklist(Seller)",        "permissionCode": "admin_logistics.whiteblacklist(seller).read"    },    {        "name": "Admin Logistics / Whitelist / Blacklist(Buyer)",        "permissionCode": "admin_logistics.whiteblacklist(buyer).read"    },    {        "name": "Admin Logistics / Warehouse Whitelist",        "permissionCode": "admin_logistics.warehouse_whitelist.read"    },    {        "name": "Admin Logistics / Update Free Shipping",        "permissionCode": "admin_logistics.update_free_shipping.read"    },    {        "name": "Admin Logistics / CB Shop decouple",        "permissionCode": "admin_logistics.configuration.read"    },    {        "name": "Admin Logistics / Ship to Location(shop)",        "permissionCode": "admin_logistics.shipto_loacation(shop).read"    },    {        "name": "Entity Management / Shop COD Blacklist",        "permissionCode": "entity_management.white/black_list.read"    },    {        "name": "3PL service management / Serviceable Area",        "permissionCode": "3pl_service_management.serviceable_area.read"    },    {        "name": "3PL service management / Dangerous Goods",        "permissionCode": "3pl_service_management.dangerous_goods.read"    },    {        "name": "Routing Forecasting / Forecasting Task",        "permissionCode": "routing_forecasting.forecasting_task.read"    },    {        "name": "Routing Forecasting / Shipment Summary",        "permissionCode": "routing_forecasting.shipment_summary.read"    },    {        "name": "Routing Forecasting / Receiver Zone(Forecasting)",        "permissionCode": "routing_forecasting.receiver_zone.read"    },    {        "name": "Basic Data Management / Reroute Info",        "permissionCode": "basic_data_management.reroute_info.read"    },    {        "name": "Basic Data Management / Transit Whs Info",        "permissionCode": "basic_data_management.warehouse_info.read"    }]`
	type dat struct {
		Name string `json:"name"`
		Code string `json:"permissionCode"`
	}
	data := make([]dat, 0)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		log.Fatalf("unmarshal failed, err=%s", err)
	}

	rows := [][]string{
		{"name", "code"},
	}
	for _, item := range data {
		rows = append(rows, []string{item.Name, item.Code})
	}
	xls, err := ExportXls(rows)
	if err != nil {
		log.Fatalf("export failed, err=%s", err)
	}

	err = xls.SaveAs("out.xls")
	if err != nil {
		log.Fatalf("save failed, err=%s", err)
	}
	// output:

}
