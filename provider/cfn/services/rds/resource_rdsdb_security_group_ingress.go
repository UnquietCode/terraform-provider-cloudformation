// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSDBSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRDSDBSecurityGroupIngressCreate,
		Read:   resourceRDSDBSecurityGroupIngressRead,
		Update: resourceRDSDBSecurityGroupIngressUpdate,
		Delete: resourceRDSDBSecurityGroupIngressDelete,

		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Required: false,
			},
			"db_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"ec2_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceRDSDBSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBSecurityGroupIngress", data, meta)
}

func resourceRDSDBSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBSecurityGroupIngress", data, meta)
}

func resourceRDSDBSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBSecurityGroupIngress", data, meta)
}

func resourceRDSDBSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBSecurityGroupIngress", data, meta)
}