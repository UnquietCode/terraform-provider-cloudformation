// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html

package cloud9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloud9EnvironmentEC2Type string = "AWS::Cloud9::EnvironmentEC2"

func ResourceCloud9EnvironmentEC2() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloud9EnvironmentEC2Read,
		Create: resourceCloud9EnvironmentEC2Create,
		Update: resourceCloud9EnvironmentEC2Update,
		Delete: resourceCloud9EnvironmentEC2Delete,
		CustomizeDiff: resourceCloud9EnvironmentEC2CustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"repositories": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentEC2Repository(),
				Optional: true,
			},
			"owner_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"automatic_stop_time_minutes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func resourceCloud9EnvironmentEC2Read(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloud9EnvironmentEC2Type, ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Create(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloud9EnvironmentEC2Type, ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Update(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloud9EnvironmentEC2Type, ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Delete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloud9EnvironmentEC2Type, data, meta)
}

func resourceCloud9EnvironmentEC2CustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloud9EnvironmentEC2Type, data, meta)
}
