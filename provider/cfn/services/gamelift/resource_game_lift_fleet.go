// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGameLiftFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameLiftFleetCreate,
		Read:   resourceGameLiftFleetRead,
		Update: resourceGameLiftFleetUpdate,
		Delete: resourceGameLiftFleetDelete,

		Schema: map[string]*schema.Schema{
			"build_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"desired_ec2_instances": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ec2_inbound_permissions": {
				Type: schema.TypeSet,
				Elem: propertyIpPermission(),
				Required: false,
			},
			"ec2_instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_paths": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"max_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"min_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"server_launch_parameters": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"server_launch_path": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGameLiftFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Fleet", data, meta)
}

func resourceGameLiftFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Fleet", data, meta)
}

func resourceGameLiftFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Fleet", data, meta)
}

func resourceGameLiftFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Fleet", data, meta)
}