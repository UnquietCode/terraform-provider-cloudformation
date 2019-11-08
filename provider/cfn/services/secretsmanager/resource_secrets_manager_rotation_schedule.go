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

func ResourceSecretsManagerRotationSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretsManagerRotationScheduleCreate,
		Read:   resourceSecretsManagerRotationScheduleRead,
		Update: resourceSecretsManagerRotationScheduleUpdate,
		Delete: resourceSecretsManagerRotationScheduleDelete,

		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rotation_lambda_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"rotation_rules": {
				Type: schema.TypeList,
				Elem: propertyRotationRules(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceSecretsManagerRotationScheduleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::RotationSchedule", data, meta)
}

func resourceSecretsManagerRotationScheduleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::RotationSchedule", data, meta)
}

func resourceSecretsManagerRotationScheduleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::RotationSchedule", data, meta)
}

func resourceSecretsManagerRotationScheduleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::RotationSchedule", data, meta)
}