// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGroupVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupVersionCreate,
		Read:   resourceGroupVersionRead,
		Update: resourceGroupVersionUpdate,
		Delete: resourceGroupVersionDelete,

		Schema: map[string]*schema.Schema{
			"logger_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"device_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"function_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"core_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"resource_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"connector_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"subscription_definition_version_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGroupVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::GroupVersion", data, meta)
}

func resourceGroupVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::GroupVersion", data, meta)
}

func resourceGroupVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::GroupVersion", data, meta)
}

func resourceGroupVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::GroupVersion", data, meta)
}