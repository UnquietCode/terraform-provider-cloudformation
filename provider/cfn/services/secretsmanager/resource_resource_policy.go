// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourcePolicyCreate,
		Read:   resourceResourcePolicyRead,
		Update: resourceResourcePolicyUpdate,
		Delete: resourceResourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_policy": {
				Type: schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceResourcePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceResourcePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceResourcePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceResourcePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::ResourcePolicy", data, meta)
}