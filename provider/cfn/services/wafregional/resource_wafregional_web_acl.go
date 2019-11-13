// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalWebACL() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalWebACLExists,
		Read:   resourceWAFRegionalWebACLRead,
		Create: resourceWAFRegionalWebACLCreate,
		Update: resourceWAFRegionalWebACLUpdate,
		Delete: resourceWAFRegionalWebACLDelete,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
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
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalWebACLExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalWebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::WebACL", ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::WebACL", ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::WebACL", ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::WebACL", data, meta)
}