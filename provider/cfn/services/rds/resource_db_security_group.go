// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDBSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBSecurityGroupCreate,
		Read:   resourceDBSecurityGroupRead,
		Update: resourceDBSecurityGroupUpdate,
		Delete: resourceDBSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"db_security_group_ingress": {
				Type: schema.TypeSet,
				Elem: propertyDBSecurityGroupIngress(),
				Required: true,
			},
			"ec2_vpc_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"group_description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDBSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBSecurityGroup", data, meta)
}

func resourceDBSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBSecurityGroup", data, meta)
}

func resourceDBSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBSecurityGroup", data, meta)
}

func resourceDBSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBSecurityGroup", data, meta)
}