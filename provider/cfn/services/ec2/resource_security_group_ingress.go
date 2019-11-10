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

func ResourceSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupIngressCreate,
		Read:   resourceSecurityGroupIngressRead,
		Update: resourceSecurityGroupIngressUpdate,
		Delete: resourceSecurityGroupIngressDelete,

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
			"from_port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_prefix_list_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_security_group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_security_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
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

func resourceSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SecurityGroupIngress", data, meta)
}