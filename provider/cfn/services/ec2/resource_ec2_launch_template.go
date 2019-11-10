// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceEC2LaunchTemplateCreate,
		Read:   resourceEC2LaunchTemplateRead,
		Update: resourceEC2LaunchTemplateUpdate,
		Delete: resourceEC2LaunchTemplateDelete,

		Schema: map[string]*schema.Schema{
			"launch_template_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"launch_template_data": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLaunchTemplateData(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceEC2LaunchTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceEC2LaunchTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceEC2LaunchTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceEC2LaunchTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::LaunchTemplate", data, meta)
}