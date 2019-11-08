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

func ResourceEC2EIPAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2EIPAssociationCreate,
		Read:   resourceEC2EIPAssociationRead,
		Update: resourceEC2EIPAssociationUpdate,
		Delete: resourceEC2EIPAssociationDelete,

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

func resourceEC2EIPAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEC2EIPAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEC2EIPAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIPAssociation", data, meta)
}

func resourceEC2EIPAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIPAssociation", data, meta)
}