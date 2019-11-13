// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SecurityGroupEgress() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2SecurityGroupEgressExists,
		Read:   resourceEC2SecurityGroupEgressRead,
		Create: resourceEC2SecurityGroupEgressCreate,
		Update: resourceEC2SecurityGroupEgressUpdate,
		Delete: resourceEC2SecurityGroupEgressDelete,
		
		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_prefix_list_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_security_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"to_port": {
				Type: schema.TypeInt,
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

func resourceEC2SecurityGroupEgressExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2SecurityGroupEgressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SecurityGroupEgress", ResourceEC2SecurityGroupEgress(), data, meta)
}

func resourceEC2SecurityGroupEgressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SecurityGroupEgress", ResourceEC2SecurityGroupEgress(), data, meta)
}

func resourceEC2SecurityGroupEgressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SecurityGroupEgress", ResourceEC2SecurityGroupEgress(), data, meta)
}

func resourceEC2SecurityGroupEgressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SecurityGroupEgress", data, meta)
}