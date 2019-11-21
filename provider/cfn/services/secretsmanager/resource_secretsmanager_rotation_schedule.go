// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html

package secretsmanager

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const secretsManagerRotationScheduleType string = "AWS::SecretsManager::RotationSchedule"

func ResourceSecretsManagerRotationSchedule() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSecretsManagerRotationScheduleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(secretsManagerRotationScheduleType, ResourceSecretsManagerRotationSchedule(), data, meta)
}

func resourceSecretsManagerRotationScheduleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(secretsManagerRotationScheduleType, data, meta)
}

func resourceSecretsManagerRotationScheduleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(secretsManagerRotationScheduleType, data, meta)
}
