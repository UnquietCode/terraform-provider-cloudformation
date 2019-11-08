// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecretsManagerResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretsManagerResourcePolicyCreate,
		Read:   resourceSecretsManagerResourcePolicyRead,
		Update: resourceSecretsManagerResourcePolicyUpdate,
		Delete: resourceSecretsManagerResourcePolicyDelete,

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

func resourceSecretsManagerResourcePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceSecretsManagerResourcePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceSecretsManagerResourcePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceSecretsManagerResourcePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::ResourcePolicy", data, meta)
}