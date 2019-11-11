// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceSageMakerWorkteam() *schema.Resource {
	return &schema.Resource{
		Create: resourceSageMakerWorkteamCreate,
		Read:   resourceSageMakerWorkteamRead,
		Update: resourceSageMakerWorkteamUpdate,
		Delete: resourceSageMakerWorkteamDelete,

		Schema: map[string]*schema.Schema{
			"workteam_name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"member_definitions": {
				Type: schema.TypeList,
				Elem: propertyWorkteamMemberDefinition(),
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceSageMakerWorkteamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::Workteam", ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::Workteam", ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::Workteam", ResourceSageMakerWorkteam(), data, meta)
}

func resourceSageMakerWorkteamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::Workteam", data, meta)
}