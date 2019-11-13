// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2EIPAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2EIPAssociationExists,
		Read:   resourceEC2EIPAssociationRead,
		Create: resourceEC2EIPAssociationCreate,
		Update: resourceEC2EIPAssociationUpdate,
		Delete: resourceEC2EIPAssociationDelete,
		
		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"eip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_ip_address": {
				Type: schema.TypeString,
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

func resourceEC2EIPAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2EIPAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIPAssociation", ResourceEC2EIPAssociation(), data, meta)
}

func resourceEC2EIPAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIPAssociation", ResourceEC2EIPAssociation(), data, meta)
}

func resourceEC2EIPAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIPAssociation", ResourceEC2EIPAssociation(), data, meta)
}

func resourceEC2EIPAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIPAssociation", data, meta)
}