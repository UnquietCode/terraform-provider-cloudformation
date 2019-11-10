// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secret.html

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecretsManagerSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretsManagerSecretCreate,
		Read:   resourceSecretsManagerSecretRead,
		Update: resourceSecretsManagerSecretUpdate,
		Delete: resourceSecretsManagerSecretDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"secret_string": {
				Type: schema.TypeString,
				Required: false,
			},
			"generate_secret_string": {
				Type: schema.TypeList,
				Elem: propertySecretGenerateSecretString(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSecretsManagerSecretCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::Secret", data, meta)
}

func resourceSecretsManagerSecretRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::Secret", data, meta)
}

func resourceSecretsManagerSecretUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::Secret", data, meta)
}

func resourceSecretsManagerSecretDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::Secret", data, meta)
}