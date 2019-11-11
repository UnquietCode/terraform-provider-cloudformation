// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecretsManagerResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSecretsManagerResourcePolicyExists,
		Read: resourceSecretsManagerResourcePolicyRead,
		Create: resourceSecretsManagerResourcePolicyCreate,
		Update: resourceSecretsManagerResourcePolicyUpdate,
		Delete: resourceSecretsManagerResourcePolicyDelete,
		CustomizeDiff: resourceSecretsManagerResourcePolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_policy": {
				Type: schema.TypeMap,
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

func resourceSecretsManagerResourcePolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSecretsManagerResourcePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::ResourcePolicy", ResourceSecretsManagerResourcePolicy(), data, meta)
}

func resourceSecretsManagerResourcePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::ResourcePolicy", ResourceSecretsManagerResourcePolicy(), data, meta)
}

func resourceSecretsManagerResourcePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::ResourcePolicy", ResourceSecretsManagerResourcePolicy(), data, meta)
}

func resourceSecretsManagerResourcePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::ResourcePolicy", data, meta)
}

func resourceSecretsManagerResourcePolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::SecretsManager::ResourcePolicy", data, meta)
}

