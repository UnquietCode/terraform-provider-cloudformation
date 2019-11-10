// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBaiduChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceBaiduChannelCreate,
		Read:   resourceBaiduChannelRead,
		Update: resourceBaiduChannelUpdate,
		Delete: resourceBaiduChannelDelete,

		Schema: map[string]*schema.Schema{
			"secret_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBaiduChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::BaiduChannel", data, meta)
}

func resourceBaiduChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::BaiduChannel", data, meta)
}

func resourceBaiduChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::BaiduChannel", data, meta)
}

func resourceBaiduChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::BaiduChannel", data, meta)
}