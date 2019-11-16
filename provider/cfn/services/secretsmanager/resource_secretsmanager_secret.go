// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const secretsManagerSecretType string = "AWS::SecretsManager::Secret"

var secretsManagerSecretProperties map[string]string = map[string]string{
	"description": "Description",
	"kms_key_id": "KmsKeyId",
	"secret_string": "SecretString",
	"generate_secret_string": "GenerateSecretString",
	"tags": "Tags",
	"name": "Name",
}

func ResourceSecretsManagerSecret() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSecretsManagerSecretExists,
		Read: resourceSecretsManagerSecretRead,
		Create: resourceSecretsManagerSecretCreate,
		Update: resourceSecretsManagerSecretUpdate,
		Delete: resourceSecretsManagerSecretDelete,
		CustomizeDiff: resourceSecretsManagerSecretCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"secret_string": {
				Type: schema.TypeString,
				Optional: true,
			},
			"generate_secret_string": {
				Type: schema.TypeList,
				Elem: propertySecretGenerateSecretString(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSecretsManagerSecretExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSecretsManagerSecretRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, meta)
}

func resourceSecretsManagerSecretCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, secretsManagerSecretProperties, meta)
}

func resourceSecretsManagerSecretUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, secretsManagerSecretProperties, meta)
}

func resourceSecretsManagerSecretDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(secretsManagerSecretType, data, meta)
}

func resourceSecretsManagerSecretCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(secretsManagerSecretType, data, meta)
}
