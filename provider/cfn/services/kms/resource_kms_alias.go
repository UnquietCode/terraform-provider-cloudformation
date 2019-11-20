// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html

package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kMSAliasType string = "AWS::KMS::Alias"

func ResourceKMSAlias() *schema.Resource {
	return &schema.Resource{
		Read: resourceKMSAliasRead,
		Create: resourceKMSAliasCreate,
		Update: resourceKMSAliasUpdate,
		Delete: resourceKMSAliasDelete,
		CustomizeDiff: resourceKMSAliasCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"alias_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_key_id": {
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

func resourceKMSAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kMSAliasType, ResourceKMSAlias(), data, meta)
}

func resourceKMSAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kMSAliasType, ResourceKMSAlias(), data, meta)
}

func resourceKMSAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kMSAliasType, ResourceKMSAlias(), data, meta)
}

func resourceKMSAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kMSAliasType, data, meta)
}

func resourceKMSAliasCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kMSAliasType, data, meta)
}
