// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SecurityGroupIngressCreate,
		Read:   resourceEC2SecurityGroupIngressRead,
		Update: resourceEC2SecurityGroupIngressUpdate,
		Delete: resourceEC2SecurityGroupIngressDelete,

		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_port": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_prefix_list_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_security_group_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_security_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_security_group_owner_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"to_port": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SecurityGroupIngress", ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SecurityGroupIngress", ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SecurityGroupIngress", ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SecurityGroupIngress", data, meta)
}