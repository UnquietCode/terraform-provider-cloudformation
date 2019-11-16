// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const secretsManagerRotationScheduleType string = "AWS::SecretsManager::RotationSchedule"

var secretsManagerRotationScheduleProperties map[string]string = map[string]string{
	"secret_id": "SecretId",
	"rotation_lambda_arn": "RotationLambdaARN",
	"rotation_rules": "RotationRules",
}

func ResourceSecretsManagerRotationSchedule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSecretsManagerRotationScheduleExists,
		Read: resourceSecretsManagerRotationScheduleRead,
		Create: resourceSecretsManagerRotationScheduleCreate,
		Update: resourceSecretsManagerRotationScheduleUpdate,
		Delete: resourceSecretsManagerRotationScheduleDelete,
		CustomizeDiff: resourceSecretsManagerRotationScheduleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"rotation_lambda_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rotation_rules": {
				Type: schema.TypeSet,
				Elem: propertyRotationScheduleRotationRules(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSecretsManagerRotationScheduleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSecretsManagerRotationScheduleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, secretsManagerRotationScheduleProperties, meta)
}

func resourceSecretsManagerRotationScheduleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, secretsManagerRotationScheduleProperties, meta)
}

func resourceSecretsManagerRotationScheduleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(secretsManagerRotationScheduleType, data, meta)
}

func resourceSecretsManagerRotationScheduleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(secretsManagerRotationScheduleType, data, meta)
}
