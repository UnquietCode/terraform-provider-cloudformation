// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
				Type: schema.TypeList,
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
	return plugin.ResourceRead("AWS::SecretsManager::RotationSchedule", ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::RotationSchedule", ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::RotationSchedule", ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::RotationSchedule", data, meta)
}

func resourceSecretsManagerRotationScheduleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
