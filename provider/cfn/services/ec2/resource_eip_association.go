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

func ResourceEIPAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEIPAssociationCreate,
		Read:   resourceEIPAssociationRead,
		Update: resourceEIPAssociationUpdate,
		Delete: resourceEIPAssociationDelete,

		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"eip": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceEIPAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEIPAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEIPAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEIPAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIPAssociation", data, meta)
}