// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFWebACLType string = "AWS::WAF::WebACL"

func ResourceWAFWebACL() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFWebACLRead,
		Create: resourceWAFWebACLCreate,
		Update: resourceWAFWebACLUpdate,
		Delete: resourceWAFWebACLDelete,
		CustomizeDiff: resourceWAFWebACLCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"default_action": {
				Type: schema.TypeSet,
				Elem: propertyWebACLWafAction(),
				Required: true,
				MaxItems: 1,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rules": {
				Type: schema.TypeSet,
				Elem: propertyWebACLActivatedRule(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceWAFWebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFWebACLType, ResourceWAFWebACL(), data, meta)
}

func resourceWAFWebACLCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFWebACLType, ResourceWAFWebACL(), data, meta)
}

func resourceWAFWebACLUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFWebACLType, ResourceWAFWebACL(), data, meta)
}

func resourceWAFWebACLDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFWebACLType, data, meta)
}

func resourceWAFWebACLCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFWebACLType, data, meta)
}
