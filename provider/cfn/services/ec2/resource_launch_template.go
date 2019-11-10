// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLaunchTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceLaunchTemplateCreate,
		Read:   resourceLaunchTemplateRead,
		Update: resourceLaunchTemplateUpdate,
		Delete: resourceLaunchTemplateDelete,

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

func resourceLaunchTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceLaunchTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceLaunchTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::LaunchTemplate", data, meta)
}

func resourceLaunchTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::LaunchTemplate", data, meta)
}