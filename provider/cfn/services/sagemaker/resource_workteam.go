// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWorkteam() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkteamCreate,
		Read:   resourceWorkteamRead,
		Update: resourceWorkteamUpdate,
		Delete: resourceWorkteamDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_configuration": {
				Type: schema.TypeList,
				Elem: propertyWorkteamNotificationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"workteam_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"member_definitions": {
				Type: schema.TypeList,
				Elem: propertyWorkteamMemberDefinition(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceWorkteamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::Workteam", data, meta)
}

func resourceWorkteamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::Workteam", data, meta)
}

func resourceWorkteamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::Workteam", data, meta)
}

func resourceWorkteamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::Workteam", data, meta)
}