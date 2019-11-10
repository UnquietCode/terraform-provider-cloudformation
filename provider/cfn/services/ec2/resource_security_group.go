// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupCreate,
		Read:   resourceSecurityGroupRead,
		Update: resourceSecurityGroupUpdate,
		Delete: resourceSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_group_egress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupEgress(),
				Required: false,
			},
			"security_group_ingress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupIngress(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SecurityGroup", data, meta)
}

func resourceSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SecurityGroup", data, meta)
}

func resourceSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SecurityGroup", data, meta)
}

func resourceSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SecurityGroup", data, meta)
}