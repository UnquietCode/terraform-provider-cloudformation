// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html

package cloud9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloud9EnvironmentEC2() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloud9EnvironmentEC2Create,
		Read:   resourceCloud9EnvironmentEC2Read,
		Update: resourceCloud9EnvironmentEC2Update,
		Delete: resourceCloud9EnvironmentEC2Delete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"repositories": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentEC2Repository(),
				Optional: true,
				ForceNew: true,
			},
			"owner_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"automatic_stop_time_minutes": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloud9EnvironmentEC2Create(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cloud9::EnvironmentEC2", ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Read(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cloud9::EnvironmentEC2", ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Update(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cloud9::EnvironmentEC2", ResourceCloud9EnvironmentEC2(), data, meta)
}

func resourceCloud9EnvironmentEC2Delete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cloud9::EnvironmentEC2", data, meta)
}