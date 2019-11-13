// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2LaunchTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2LaunchTemplateExists,
		Read:   resourceEC2LaunchTemplateRead,
		Create: resourceEC2LaunchTemplateCreate,
		Update: resourceEC2LaunchTemplateUpdate,
		Delete: resourceEC2LaunchTemplateDelete,
		CustomizeDiff: resourceEC2LaunchTemplateCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"launch_template_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_template_data": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLaunchTemplateData(),
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

func resourceEC2LaunchTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2LaunchTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::LaunchTemplate", ResourceEC2LaunchTemplate(), data, meta)
}

func resourceEC2LaunchTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::LaunchTemplate", ResourceEC2LaunchTemplate(), data, meta)
}

func resourceEC2LaunchTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::LaunchTemplate", ResourceEC2LaunchTemplate(), data, meta)
}

func resourceEC2LaunchTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceEC2LaunchTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}