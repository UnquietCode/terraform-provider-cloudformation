// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFWebACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFWebACLCreate,
		Read:   resourceWAFWebACLRead,
		Update: resourceWAFWebACLUpdate,
		Delete: resourceWAFWebACLDelete,

		Schema: map[string]*schema.Schema{
			"default_action": {
				Type: schema.TypeList,
				Elem: propertyWafAction(),
				Required: true,
				MaxItems: 1,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rules": {
				Type: schema.TypeSet,
				Elem: propertyActivatedRule(),
				Required: false,
			},
		},
	}
}

func resourceWAFWebACLCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::WebACL", data, meta)
}

func resourceWAFWebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::WebACL", data, meta)
}

func resourceWAFWebACLUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::WebACL", data, meta)
}

func resourceWAFWebACLDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::WebACL", data, meta)
}