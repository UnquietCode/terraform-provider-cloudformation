// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFIPSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFIPSetCreate,
		Read:   resourceWAFIPSetRead,
		Update: resourceWAFIPSetUpdate,
		Delete: resourceWAFIPSetDelete,

		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeSet,
				Elem: propertyIPSetIPSetDescriptor(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::IPSet", data, meta)
}

func resourceWAFIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::IPSet", data, meta)
}

func resourceWAFIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::IPSet", data, meta)
}

func resourceWAFIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::IPSet", data, meta)
}