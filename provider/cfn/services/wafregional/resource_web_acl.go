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

func ResourceWebACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebACLCreate,
		Read:   resourceWebACLRead,
		Update: resourceWebACLUpdate,
		Delete: resourceWebACLDelete,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_action": {
				Type: schema.TypeList,
				Elem: propertyWebACLAction(),
				Required: true,
				MaxItems: 1,
			},
			"rules": {
				Type: schema.TypeList,
				Elem: propertyWebACLRule(),
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

func resourceWebACLCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::WebACL", data, meta)
}

func resourceWebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::WebACL", data, meta)
}

func resourceWebACLUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::WebACL", data, meta)
}

func resourceWebACLDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::WebACL", data, meta)
}