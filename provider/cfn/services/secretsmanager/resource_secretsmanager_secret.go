// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secret.html

package secretsmanager

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const secretsManagerSecretType string = "AWS::SecretsManager::Secret"

func ResourceSecretsManagerSecret() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
				Elem: propertySecretGenerateSecretString(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
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

func resourceSecretsManagerSecretRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, meta)
}

func resourceSecretsManagerSecretCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, meta)
}

func resourceSecretsManagerSecretUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(secretsManagerSecretType, ResourceSecretsManagerSecret(), data, meta)
}

func resourceSecretsManagerSecretDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(secretsManagerSecretType, data, meta)
}

func resourceSecretsManagerSecretCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(secretsManagerSecretType, data, meta)
}
