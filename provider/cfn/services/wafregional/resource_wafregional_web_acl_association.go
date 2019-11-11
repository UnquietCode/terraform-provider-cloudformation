// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalWebACLAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalWebACLAssociationExists,
		Read: resourceWAFRegionalWebACLAssociationRead,
		Create: resourceWAFRegionalWebACLAssociationCreate,
		Update: resourceWAFRegionalWebACLAssociationUpdate,
		Delete: resourceWAFRegionalWebACLAssociationDelete,
		CustomizeDiff: resourceWAFRegionalWebACLAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resource_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"web_acl_id": {
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

func resourceWAFRegionalWebACLAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalWebACLAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::WebACLAssociation", ResourceWAFRegionalWebACLAssociation(), data, meta)
}

func resourceWAFRegionalWebACLAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::WebACLAssociation", ResourceWAFRegionalWebACLAssociation(), data, meta)
}

func resourceWAFRegionalWebACLAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::WebACLAssociation", ResourceWAFRegionalWebACLAssociation(), data, meta)
}

func resourceWAFRegionalWebACLAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::WebACLAssociation", data, meta)
}

func resourceWAFRegionalWebACLAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::WAFRegional::WebACLAssociation", data, meta)
}

