// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointEmailDedicatedIpPool() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointEmailDedicatedIpPoolCreate,
		Read:   resourcePinpointEmailDedicatedIpPoolRead,
		Update: resourcePinpointEmailDedicatedIpPoolUpdate,
		Delete: resourcePinpointEmailDedicatedIpPoolDelete,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyTags(),
				Required: false,
			},
		},
	}
}

func resourcePinpointEmailDedicatedIpPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}

func resourcePinpointEmailDedicatedIpPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}

func resourcePinpointEmailDedicatedIpPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}

func resourcePinpointEmailDedicatedIpPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}