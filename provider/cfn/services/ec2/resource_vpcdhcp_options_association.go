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

func ResourceVPCDHCPOptionsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCDHCPOptionsAssociationCreate,
		Read:   resourceVPCDHCPOptionsAssociationRead,
		Update: resourceVPCDHCPOptionsAssociationUpdate,
		Delete: resourceVPCDHCPOptionsAssociationDelete,

		Schema: map[string]*schema.Schema{
			"dhcp_options_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCDHCPOptionsAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCDHCPOptionsAssociation", data, meta)
}

func resourceVPCDHCPOptionsAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCDHCPOptionsAssociation", data, meta)
}

func resourceVPCDHCPOptionsAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCDHCPOptionsAssociation", data, meta)
}

func resourceVPCDHCPOptionsAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCDHCPOptionsAssociation", data, meta)
}