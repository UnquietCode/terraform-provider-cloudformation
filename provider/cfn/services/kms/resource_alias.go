// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliasCreate,
		Read:   resourceAliasRead,
		Update: resourceAliasUpdate,
		Delete: resourceAliasDelete,

		Schema: map[string]*schema.Schema{
			"alias_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_key_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KMS::Alias", data, meta)
}

func resourceAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KMS::Alias", data, meta)
}

func resourceAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KMS::Alias", data, meta)
}

func resourceAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KMS::Alias", data, meta)
}