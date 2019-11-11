// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoT1ClickProject() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoT1ClickProjectExists,
		Read: resourceIoT1ClickProjectRead,
		Create: resourceIoT1ClickProjectCreate,
		Update: resourceIoT1ClickProjectUpdate,
		Delete: resourceIoT1ClickProjectDelete,
		CustomizeDiff: resourceIoT1ClickProjectCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"placement_template": {
				Type: schema.TypeList,
				Elem: propertyProjectPlacementTemplate(),
				Required: true,
				MaxItems: 1,
			},
			"project_name": {
				Type: schema.TypeString,
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

func resourceIoT1ClickProjectExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoT1ClickProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Project", ResourceIoT1ClickProject(), data, meta)
}

func resourceIoT1ClickProjectCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Project", ResourceIoT1ClickProject(), data, meta)
}

func resourceIoT1ClickProjectUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Project", ResourceIoT1ClickProject(), data, meta)
}

func resourceIoT1ClickProjectDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Project", data, meta)
}

func resourceIoT1ClickProjectCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::IoT1Click::Project", data, meta)
}

