// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerWorkteamType string = "AWS::SageMaker::Workteam"

var sageMakerWorkteamProperties map[string]string = map[string]string{
	"description": "Description",
	"notification_configuration": "NotificationConfiguration",
	"workteam_name": "WorkteamName",
	"member_definitions": "MemberDefinitions",
	"tags": "Tags",
}

func ResourceSageMakerWorkteam() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSageMakerWorkteamExists,
		Read: resourceSageMakerWorkteamRead,
		Create: resourceSageMakerWorkteamCreate,
		Update: resourceSageMakerWorkteamUpdate,
		Delete: resourceSageMakerWorkteamDelete,
		CustomizeDiff: resourceSageMakerWorkteamCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_configuration": {
				Type: schema.TypeList,
				Elem: propertyWorkteamNotificationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"workteam_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"member_definitions": {
				Type: schema.TypeList,
				Elem: propertyWorkteamMemberDefinition(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSageMakerWorkteamExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerWorkteamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, sageMakerWorkteamProperties, meta)
}

func resourceSageMakerWorkteamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, sageMakerWorkteamProperties, meta)
}

func resourceSageMakerWorkteamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerWorkteamType, data, meta)
}

func resourceSageMakerWorkteamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerWorkteamType, data, meta)
}
