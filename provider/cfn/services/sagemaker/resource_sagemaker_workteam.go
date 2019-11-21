// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerWorkteamType string = "AWS::SageMaker::Workteam"

func ResourceSageMakerWorkteam() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSageMakerWorkteamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerWorkteamType, ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerWorkteamType, data, meta)
}

func resourceSageMakerWorkteamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerWorkteamType, data, meta)
}
