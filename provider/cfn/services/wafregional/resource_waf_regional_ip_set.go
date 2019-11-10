// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalIPSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalIPSetCreate,
		Read:   resourceWAFRegionalIPSetRead,
		Update: resourceWAFRegionalIPSetUpdate,
		Delete: resourceWAFRegionalIPSetDelete,

		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeList,
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

func resourceWAFRegionalIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::IPSet", data, meta)
}

func resourceWAFRegionalIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::IPSet", data, meta)
}

func resourceWAFRegionalIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::IPSet", data, meta)
}

func resourceWAFRegionalIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::IPSet", data, meta)
}