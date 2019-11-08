// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SecurityGroupEgress() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SecurityGroupEgressCreate,
		Read:   resourceEC2SecurityGroupEgressRead,
		Update: resourceEC2SecurityGroupEgressUpdate,
		Delete: resourceEC2SecurityGroupEgressDelete,

		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"destination_prefix_list_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"destination_security_group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"from_port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"to_port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SecurityGroupEgressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SecurityGroupEgress", data, meta)
}

func resourceEC2SecurityGroupEgressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SecurityGroupEgress", data, meta)
}

func resourceEC2SecurityGroupEgressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SecurityGroupEgress", data, meta)
}

func resourceEC2SecurityGroupEgressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SecurityGroupEgress", data, meta)
}