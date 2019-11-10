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

func ResourceWebACLAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebACLAssociationCreate,
		Read:   resourceWebACLAssociationRead,
		Update: resourceWebACLAssociationUpdate,
		Delete: resourceWebACLAssociationDelete,

		Schema: map[string]*schema.Schema{
			"resource_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"web_acl_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWebACLAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::WebACLAssociation", data, meta)
}

func resourceWebACLAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::WebACLAssociation", data, meta)
}

func resourceWebACLAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::WebACLAssociation", data, meta)
}

func resourceWebACLAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::WebACLAssociation", data, meta)
}